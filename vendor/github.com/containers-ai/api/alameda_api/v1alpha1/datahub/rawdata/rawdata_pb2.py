# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/rawdata/rawdata.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import rawdata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2
from alameda_api.v1alpha1.datahub.common import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2
from alameda_api.v1alpha1.datahub.rawdata import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_rawdata_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/rawdata/rawdata.proto',
  package='containersai.alameda.v1alpha1.datahub.rawdata',
  syntax='proto3',
  serialized_options=b'ZAgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/rawdata',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n2alameda_api/v1alpha1/datahub/rawdata/rawdata.proto\x12-containersai.alameda.v1alpha1.datahub.rawdata\x1a\x31\x61lameda_api/v1alpha1/datahub/common/rawdata.proto\x1a/alameda_api/v1alpha1/datahub/common/types.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/rawdata/types.proto\"\xb9\x01\n\x0bReadRawdata\x12\x43\n\x05query\x18\x01 \x01(\x0b\x32\x34.containersai.alameda.v1alpha1.datahub.rawdata.Query\x12\x0f\n\x07\x63olumns\x18\x02 \x03(\t\x12\x43\n\x06groups\x18\x03 \x03(\x0b\x32\x33.containersai.alameda.v1alpha1.datahub.common.Group\x12\x0f\n\x07rawdata\x18\x04 \x01(\t\"\x9d\x02\n\x0cWriteRawdata\x12\x10\n\x08\x64\x61tabase\x18\x01 \x01(\t\x12\r\n\x05table\x18\x02 \x01(\t\x12\x0f\n\x07\x63olumns\x18\x03 \x03(\t\x12?\n\x04rows\x18\x04 \x03(\x0b\x32\x31.containersai.alameda.v1alpha1.datahub.common.Row\x12N\n\x0c\x63olumn_types\x18\x05 \x03(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.ColumnType\x12J\n\ndata_types\x18\x06 \x03(\x0e\x32\x36.containersai.alameda.v1alpha1.datahub.common.DataTypeBCZAgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/rawdatab\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_rawdata_dot_types__pb2.DESCRIPTOR,])




_READRAWDATA = _descriptor.Descriptor(
  name='ReadRawdata',
  full_name='containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='query', full_name='containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata.query', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='columns', full_name='containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata.columns', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='groups', full_name='containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata.groups', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='rawdata', full_name='containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata.rawdata', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=252,
  serialized_end=437,
)


_WRITERAWDATA = _descriptor.Descriptor(
  name='WriteRawdata',
  full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='database', full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.database', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='table', full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.table', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='columns', full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.columns', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='rows', full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.rows', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='column_types', full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.column_types', index=4,
      number=5, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data_types', full_name='containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.data_types', index=5,
      number=6, type=14, cpp_type=8, label=3,
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
  serialized_start=440,
  serialized_end=725,
)

_READRAWDATA.fields_by_name['query'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_rawdata_dot_types__pb2._QUERY
_READRAWDATA.fields_by_name['groups'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2._GROUP
_WRITERAWDATA.fields_by_name['rows'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2._ROW
_WRITERAWDATA.fields_by_name['column_types'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._COLUMNTYPE
_WRITERAWDATA.fields_by_name['data_types'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._DATATYPE
DESCRIPTOR.message_types_by_name['ReadRawdata'] = _READRAWDATA
DESCRIPTOR.message_types_by_name['WriteRawdata'] = _WRITERAWDATA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ReadRawdata = _reflection.GeneratedProtocolMessageType('ReadRawdata', (_message.Message,), {
  'DESCRIPTOR' : _READRAWDATA,
  '__module__' : 'alameda_api.v1alpha1.datahub.rawdata.rawdata_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata)
  })
_sym_db.RegisterMessage(ReadRawdata)

WriteRawdata = _reflection.GeneratedProtocolMessageType('WriteRawdata', (_message.Message,), {
  'DESCRIPTOR' : _WRITERAWDATA,
  '__module__' : 'alameda_api.v1alpha1.datahub.rawdata.rawdata_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata)
  })
_sym_db.RegisterMessage(WriteRawdata)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)