package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

var JobDone = errors.New("job done")

func NewJob(ty JobType, pb protobuf.Message) *Job {
	bs, err := protobuf.Marshal(pb)
	if err != nil {
		panic(fmt.Errorf("protobuf marshal: %v", err))
	}

	j := &Job{
		Id:      uuid.NewString(),
		Type:    ty,
		Content: bs,
	}
	return j
}

func NewJobFromBytes(bs []byte) *Job {
	j := &Job{}
	err := protobuf.Unmarshal(bs, j)
	if err != nil {
		panic("invalid task")
	}
	return j
}

func (d *DB) InsertJob(j *Job) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	bs, err := protobuf.Marshal(j)
	if err != nil {
		return err
	}

	err = txn.Set(JobKey(j.Id), bs)
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (d *DB) ListJobs() {
	panic("implement me")
}

func (d *DB) GetJob(ctx context.Context, jobId string) (j *Job, err error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(JobKey(jobId))
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, NewNotFoundErr(jobId)
		}
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		j = NewJobFromBytes(val)
		return nil
	})
	return
}

func (d *DB) SubscribeJob(ctx context.Context, fn func(j *Job)) (err error) {
	return d.db.Subscribe(ctx, func(kv *badger.KVList) error {
		for _, v := range kv.Kv {
			j := NewJobFromBytes(v.Value)
			fn(j)
		}
		return nil
	}, JobPrefix)
}

func (d *DB) DeleteJob(ctx context.Context, jobId string) (err error) {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	err = txn.Delete(JobKey(jobId))
	if err != nil {
		return
	}
	return
}

func (d *DB) WaitJob(ctx context.Context, jobId string) (err error) {
	_, err = d.GetJob(ctx, jobId)
	// If job doesn't exist, we can return directly.
	if err != nil && errors.Is(err, ErrNotFound) {
		return nil
	}
	if err != nil {
		return
	}

	err = d.db.Subscribe(ctx, func(kv *badger.KVList) error {
		for _, v := range kv.Kv {
			if v.Value == nil {
				return JobDone
			}
		}
		return nil
	}, JobKey(jobId))
	if err == JobDone {
		return nil
	}
	return err
}
