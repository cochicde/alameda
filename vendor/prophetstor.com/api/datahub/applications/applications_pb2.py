# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/applications/applications.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.common import queries_pb2 as datahub_dot_common_dot_queries__pb2
from datahub.common import rawdata_pb2 as datahub_dot_common_dot_rawdata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/applications/applications.proto',
  package='prophetstor.datahub.applications',
  syntax='proto3',
  serialized_options=b'Z(prophetstor.com/api/datahub/applications',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\'datahub/applications/applications.proto\x12 prophetstor.datahub.applications\x1a\x1c\x64\x61tahub/common/queries.proto\x1a\x1c\x64\x61tahub/common/rawdata.proto\"b\n\x10WriteApplication\x12\x13\n\x0bmeasurement\x18\x01 \x01(\t\x12\x39\n\nwrite_data\x18\x02 \x01(\x0b\x32%.prophetstor.datahub.common.WriteData\"f\n\x0fReadApplication\x12\x13\n\x0bmeasurement\x18\x01 \x01(\t\x12>\n\x0fwhere_condition\x18\x02 \x03(\x0b\x32%.prophetstor.datahub.common.Condition\"h\n\x11\x44\x65leteApplication\x12\x13\n\x0bmeasurement\x18\x01 \x01(\t\x12>\n\x0fwhere_condition\x18\x02 \x03(\x0b\x32%.prophetstor.datahub.common.ConditionB*Z(prophetstor.com/api/datahub/applicationsb\x06proto3'
  ,
  dependencies=[datahub_dot_common_dot_queries__pb2.DESCRIPTOR,datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,])




_WRITEAPPLICATION = _descriptor.Descriptor(
  name='WriteApplication',
  full_name='prophetstor.datahub.applications.WriteApplication',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='measurement', full_name='prophetstor.datahub.applications.WriteApplication.measurement', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_data', full_name='prophetstor.datahub.applications.WriteApplication.write_data', index=1,
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
  serialized_start=137,
  serialized_end=235,
)


_READAPPLICATION = _descriptor.Descriptor(
  name='ReadApplication',
  full_name='prophetstor.datahub.applications.ReadApplication',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='measurement', full_name='prophetstor.datahub.applications.ReadApplication.measurement', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='where_condition', full_name='prophetstor.datahub.applications.ReadApplication.where_condition', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=237,
  serialized_end=339,
)


_DELETEAPPLICATION = _descriptor.Descriptor(
  name='DeleteApplication',
  full_name='prophetstor.datahub.applications.DeleteApplication',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='measurement', full_name='prophetstor.datahub.applications.DeleteApplication.measurement', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='where_condition', full_name='prophetstor.datahub.applications.DeleteApplication.where_condition', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=341,
  serialized_end=445,
)

_WRITEAPPLICATION.fields_by_name['write_data'].message_type = datahub_dot_common_dot_rawdata__pb2._WRITEDATA
_READAPPLICATION.fields_by_name['where_condition'].message_type = datahub_dot_common_dot_queries__pb2._CONDITION
_DELETEAPPLICATION.fields_by_name['where_condition'].message_type = datahub_dot_common_dot_queries__pb2._CONDITION
DESCRIPTOR.message_types_by_name['WriteApplication'] = _WRITEAPPLICATION
DESCRIPTOR.message_types_by_name['ReadApplication'] = _READAPPLICATION
DESCRIPTOR.message_types_by_name['DeleteApplication'] = _DELETEAPPLICATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WriteApplication = _reflection.GeneratedProtocolMessageType('WriteApplication', (_message.Message,), {
  'DESCRIPTOR' : _WRITEAPPLICATION,
  '__module__' : 'datahub.applications.applications_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.applications.WriteApplication)
  })
_sym_db.RegisterMessage(WriteApplication)

ReadApplication = _reflection.GeneratedProtocolMessageType('ReadApplication', (_message.Message,), {
  'DESCRIPTOR' : _READAPPLICATION,
  '__module__' : 'datahub.applications.applications_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.applications.ReadApplication)
  })
_sym_db.RegisterMessage(ReadApplication)

DeleteApplication = _reflection.GeneratedProtocolMessageType('DeleteApplication', (_message.Message,), {
  'DESCRIPTOR' : _DELETEAPPLICATION,
  '__module__' : 'datahub.applications.applications_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.applications.DeleteApplication)
  })
_sym_db.RegisterMessage(DeleteApplication)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)