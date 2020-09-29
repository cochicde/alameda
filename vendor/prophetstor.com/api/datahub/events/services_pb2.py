# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/events/services.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.common import queries_pb2 as datahub_dot_common_dot_queries__pb2
from datahub.events import events_pb2 as datahub_dot_events_dot_events__pb2
from datahub.events import types_pb2 as datahub_dot_events_dot_types__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/events/services.proto',
  package='prophetstor.datahub.events',
  syntax='proto3',
  serialized_options=b'Z\"prophetstor.com/api/datahub/events',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1d\x64\x61tahub/events/services.proto\x12\x1aprophetstor.datahub.events\x1a\x1c\x64\x61tahub/common/queries.proto\x1a\x1b\x64\x61tahub/events/events.proto\x1a\x1a\x64\x61tahub/events/types.proto\x1a\x17google/rpc/status.proto\"\x9f\x02\n\x11ListEventsRequest\x12\x43\n\x0fquery_condition\x18\x01 \x01(\x0b\x32*.prophetstor.datahub.common.QueryCondition\x12\n\n\x02id\x18\x02 \x03(\t\x12\x12\n\ncluster_id\x18\x03 \x03(\t\x12\x33\n\x04type\x18\x04 \x03(\x0e\x32%.prophetstor.datahub.events.EventType\x12\x39\n\x07version\x18\x05 \x03(\x0e\x32(.prophetstor.datahub.events.EventVersion\x12\x35\n\x05level\x18\x06 \x03(\x0e\x32&.prophetstor.datahub.events.EventLevel\"k\n\x12ListEventsResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x31\n\x06\x65vents\x18\x02 \x03(\x0b\x32!.prophetstor.datahub.events.Event\"H\n\x13\x43reateEventsRequest\x12\x31\n\x06\x65vents\x18\x01 \x03(\x0b\x32!.prophetstor.datahub.events.EventB$Z\"prophetstor.com/api/datahub/eventsb\x06proto3'
  ,
  dependencies=[datahub_dot_common_dot_queries__pb2.DESCRIPTOR,datahub_dot_events_dot_events__pb2.DESCRIPTOR,datahub_dot_events_dot_types__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_LISTEVENTSREQUEST = _descriptor.Descriptor(
  name='ListEventsRequest',
  full_name='prophetstor.datahub.events.ListEventsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='prophetstor.datahub.events.ListEventsRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='prophetstor.datahub.events.ListEventsRequest.id', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cluster_id', full_name='prophetstor.datahub.events.ListEventsRequest.cluster_id', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='prophetstor.datahub.events.ListEventsRequest.type', index=3,
      number=4, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='prophetstor.datahub.events.ListEventsRequest.version', index=4,
      number=5, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='level', full_name='prophetstor.datahub.events.ListEventsRequest.level', index=5,
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
  serialized_start=174,
  serialized_end=461,
)


_LISTEVENTSRESPONSE = _descriptor.Descriptor(
  name='ListEventsResponse',
  full_name='prophetstor.datahub.events.ListEventsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='prophetstor.datahub.events.ListEventsResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='events', full_name='prophetstor.datahub.events.ListEventsResponse.events', index=1,
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
  serialized_start=463,
  serialized_end=570,
)


_CREATEEVENTSREQUEST = _descriptor.Descriptor(
  name='CreateEventsRequest',
  full_name='prophetstor.datahub.events.CreateEventsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='events', full_name='prophetstor.datahub.events.CreateEventsRequest.events', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=572,
  serialized_end=644,
)

_LISTEVENTSREQUEST.fields_by_name['query_condition'].message_type = datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_LISTEVENTSREQUEST.fields_by_name['type'].enum_type = datahub_dot_events_dot_types__pb2._EVENTTYPE
_LISTEVENTSREQUEST.fields_by_name['version'].enum_type = datahub_dot_events_dot_types__pb2._EVENTVERSION
_LISTEVENTSREQUEST.fields_by_name['level'].enum_type = datahub_dot_events_dot_types__pb2._EVENTLEVEL
_LISTEVENTSRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTEVENTSRESPONSE.fields_by_name['events'].message_type = datahub_dot_events_dot_events__pb2._EVENT
_CREATEEVENTSREQUEST.fields_by_name['events'].message_type = datahub_dot_events_dot_events__pb2._EVENT
DESCRIPTOR.message_types_by_name['ListEventsRequest'] = _LISTEVENTSREQUEST
DESCRIPTOR.message_types_by_name['ListEventsResponse'] = _LISTEVENTSRESPONSE
DESCRIPTOR.message_types_by_name['CreateEventsRequest'] = _CREATEEVENTSREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListEventsRequest = _reflection.GeneratedProtocolMessageType('ListEventsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTEVENTSREQUEST,
  '__module__' : 'datahub.events.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.events.ListEventsRequest)
  })
_sym_db.RegisterMessage(ListEventsRequest)

ListEventsResponse = _reflection.GeneratedProtocolMessageType('ListEventsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTEVENTSRESPONSE,
  '__module__' : 'datahub.events.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.events.ListEventsResponse)
  })
_sym_db.RegisterMessage(ListEventsResponse)

CreateEventsRequest = _reflection.GeneratedProtocolMessageType('CreateEventsRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEEVENTSREQUEST,
  '__module__' : 'datahub.events.services_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.events.CreateEventsRequest)
  })
_sym_db.RegisterMessage(CreateEventsRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)