# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/scores/scores.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/scores/scores.proto',
  package='prophetstor.datahub.scores',
  syntax='proto3',
  serialized_options=b'Z\"prophetstor.com/api/datahub/scores',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1b\x64\x61tahub/scores/scores.proto\x12\x1aprophetstor.datahub.scores\x1a\x1fgoogle/protobuf/timestamp.proto\"o\n\x18SimulatedSchedulingScore\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x14\n\x0cscore_before\x18\x02 \x01(\x01\x12\x13\n\x0bscore_after\x18\x03 \x01(\x01\x42$Z\"prophetstor.com/api/datahub/scoresb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_SIMULATEDSCHEDULINGSCORE = _descriptor.Descriptor(
  name='SimulatedSchedulingScore',
  full_name='prophetstor.datahub.scores.SimulatedSchedulingScore',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='prophetstor.datahub.scores.SimulatedSchedulingScore.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='score_before', full_name='prophetstor.datahub.scores.SimulatedSchedulingScore.score_before', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='score_after', full_name='prophetstor.datahub.scores.SimulatedSchedulingScore.score_after', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=92,
  serialized_end=203,
)

_SIMULATEDSCHEDULINGSCORE.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['SimulatedSchedulingScore'] = _SIMULATEDSCHEDULINGSCORE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SimulatedSchedulingScore = _reflection.GeneratedProtocolMessageType('SimulatedSchedulingScore', (_message.Message,), {
  'DESCRIPTOR' : _SIMULATEDSCHEDULINGSCORE,
  '__module__' : 'datahub.scores.scores_pb2'
  # @@protoc_insertion_point(class_scope:prophetstor.datahub.scores.SimulatedSchedulingScore)
  })
_sym_db.RegisterMessage(SimulatedSchedulingScore)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)