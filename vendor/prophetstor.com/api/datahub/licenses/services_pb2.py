# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/licenses/services.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.licenses import licenses_pb2 as datahub_dot_licenses_dot_licenses__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/licenses/services.proto',
  package='prophetstor.datahub.licenses',
  syntax='proto3',
  serialized_options=b'Z$prophetstor.com/api/datahub/licenses',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1f\x64\x61tahub/licenses/services.proto\x12\x1cprophetstor.datahub.licenses\x1a\x1f\x64\x61tahub/licenses/licenses.proto\x1a\x17google/rpc/status.proto\"p\n\x12GetLicenseResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x36\n\x07license\x18\x02 \x01(\x0b\x32%.prophetstor.datahub.licenses.LicenseB&Z$prophetstor.com/api/datahub/licensesb\x06proto3'
  ,
  dependencies=[datahub_dot_licenses_dot_licenses__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_GETLICENSERESPONSE = _descriptor.Descriptor(
  name='GetLicenseResponse',
  full_name='prophetstor.datahub.licenses.GetLicenseResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.datahub.licenses.GetLicenseResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='license', full_name='prophetstor.datahub.licenses.GetLicenseResponse.license', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=123,
  serialized_end=235,
)

_GETLICENSERESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_GETLICENSERESPONSE.fields_by_name['license'].message_type = datahub_dot_licenses_dot_licenses__pb2._LICENSE
DESCRIPTOR.message_types_by_name['GetLicenseResponse'] = _GETLICENSERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetLicenseResponse = _reflection.GeneratedProtocolMessageType('GetLicenseResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETLICENSERESPONSE,
  '__module__' : 'datahub.licenses.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.licenses.GetLicenseResponse)
  })
_sym_db.RegisterMessage(GetLicenseResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
