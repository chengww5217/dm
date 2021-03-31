import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import './controller.dart';

class DestinationFormFields extends GetView<CreateTaskController> {
  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: EdgeInsets.only(top: 16.0, bottom: 32.0),
          child: SelectableText(
            'Destination set'.tr,
            style: TextStyle(
              fontSize: 12,
              fontWeight: FontWeight.w600,
              height: 1.67,
            ),
          ),
        ),
        Container(
          padding: EdgeInsets.symmetric(horizontal: 30.0),
          margin: EdgeInsets.only(bottom: 24.0),
          child: Row(
            children: [
              SizedBox(
                width: 124,
                child: RichText(
                  text: TextSpan(
                    text: 'Destination type'.tr,
                    style: TextStyle(
                      fontSize: 12,
                      fontWeight: FontWeight.w500,
                    ),
                    children: [
                      TextSpan(
                        text: '*',
                        style: TextStyle(
                          color: Colors.red,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Expanded(
                child: DropdownButtonFormField(
                  hint: Text('Select destination type'.tr),
                  isExpanded: true,
                  isDense: true,
                  style: TextStyle(fontSize: 12),
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                  ),
                  value: controller.dstType.value != ''
                      ? controller.dstType.value
                      : null,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please select destination type')
                      .build(),
                  onChanged: controller.dstType,
                  onSaved: controller.dstType,
                  items: [
                    DropdownMenuItem(
                      child: Text('QingStor'.tr),
                      value: 'qingstor',
                    ),
                    DropdownMenuItem(
                      child: Text('FS'.tr),
                      value: 'fs',
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
        Container(
          padding: EdgeInsets.symmetric(horizontal: 30.0),
          margin: EdgeInsets.only(bottom: 32.0),
          child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
            SizedBox(
              width: 124,
              child: RichText(
                text: TextSpan(
                  text: 'Destination path'.tr,
                  style: TextStyle(
                    fontSize: 12,
                    height: 2.67,
                    fontWeight: FontWeight.w500,
                  ),
                  children: [
                    TextSpan(
                      text: '*',
                      style: TextStyle(
                        color: Colors.red,
                      ),
                    ),
                  ],
                ),
              ),
            ),
            Expanded(
              child: TextFormField(
                decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                ),
                maxLines: 1,
                style: TextStyle(fontSize: 12),
                textInputAction: TextInputAction.next,
                keyboardType: TextInputType.text,
                initialValue: '/',
                validator: ValidationBuilder()
                    .minLength(1, 'Please enter destination path')
                    .build(),
                onSaved: controller.dstPath,
              ),
            ),
          ]),
        ),
        Obx(
          () => Visibility(
            visible: controller.dstType.value == 'qingstor',
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 30.0),
              margin: EdgeInsets.only(bottom: 32.0),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SizedBox(
                    width: 124,
                    child: RichText(
                      text: TextSpan(
                        text: 'Bucket name'.tr,
                        style: TextStyle(
                          fontSize: 12,
                          height: 2.67,
                          fontWeight: FontWeight.w500,
                        ),
                        children: [
                          TextSpan(
                            text: '*',
                            style: TextStyle(
                              color: Colors.red,
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                  Expanded(
                    child: TextFormField(
                      decoration: InputDecoration(
                        border: const OutlineInputBorder(),
                        contentPadding:
                            EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                      ),
                      maxLines: 1,
                      style: TextStyle(fontSize: 12),
                      textInputAction: TextInputAction.next,
                      keyboardType: TextInputType.text,
                      validator: ValidationBuilder()
                          .minLength(1, 'Please enter bucket name')
                          .build(),
                      onSaved: controller.dstBucketName,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
        Obx(
          () => Visibility(
            visible: controller.dstType.value == 'qingstor',
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 30.0),
              margin: EdgeInsets.only(bottom: 32.0),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SizedBox(
                    width: 124,
                    child: RichText(
                      text: TextSpan(
                        text: 'Credential'.tr,
                        style: TextStyle(
                          fontSize: 12,
                          height: 2.67,
                          fontWeight: FontWeight.w500,
                        ),
                        children: [
                          TextSpan(
                            text: '*',
                            style: TextStyle(
                              color: Colors.red,
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                  Expanded(
                    child: TextFormField(
                      decoration: InputDecoration(
                        border: const OutlineInputBorder(),
                        contentPadding:
                            EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                      ),
                      maxLines: 1,
                      style: TextStyle(fontSize: 12),
                      textInputAction: TextInputAction.next,
                      keyboardType: TextInputType.text,
                      validator: ValidationBuilder()
                          .minLength(1, 'Please enter credential')
                          .build(),
                      onSaved: controller.dstCredential,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ],
    );
  }
}