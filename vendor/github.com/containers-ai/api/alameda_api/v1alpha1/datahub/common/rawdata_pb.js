// source: alameda_api/v1alpha1/datahub/common/rawdata.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.Group', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.ReadData', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.Row', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.WriteData', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.common.WriteData.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.common.WriteData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.common.WriteData.displayName = 'proto.containersai.alameda.v1alpha1.datahub.common.WriteData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.common.ReadData.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.common.ReadData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.common.ReadData.displayName = 'proto.containersai.alameda.v1alpha1.datahub.common.ReadData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.common.Row.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.common.Row, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.common.Row.displayName = 'proto.containersai.alameda.v1alpha1.datahub.common.Row';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.common.Group.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.common.Group, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.common.Group.displayName = 'proto.containersai.alameda.v1alpha1.datahub.common.Group';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.common.WriteData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.toObject = function(includeInstance, msg) {
  var f, obj = {
    columnsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    rowsList: jspb.Message.toObjectList(msg.getRowsList(),
    proto.containersai.alameda.v1alpha1.datahub.common.Row.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.common.WriteData;
  return proto.containersai.alameda.v1alpha1.datahub.common.WriteData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addColumns(value);
      break;
    case 2:
      var value = new proto.containersai.alameda.v1alpha1.datahub.common.Row;
      reader.readMessage(value,proto.containersai.alameda.v1alpha1.datahub.common.Row.deserializeBinaryFromReader);
      msg.addRows(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.common.WriteData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getColumnsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getRowsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.containersai.alameda.v1alpha1.datahub.common.Row.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string columns = 1;
 * @return {!Array<string>}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.getColumnsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.setColumnsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.addColumns = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.clearColumnsList = function() {
  return this.setColumnsList([]);
};


/**
 * repeated Row rows = 2;
 * @return {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Row>}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.getRowsList = function() {
  return /** @type{!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Row>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.containersai.alameda.v1alpha1.datahub.common.Row, 2));
};


/**
 * @param {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Row>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.setRowsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Row=} opt_value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row}
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.addRows = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.containersai.alameda.v1alpha1.datahub.common.Row, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.WriteData} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.WriteData.prototype.clearRowsList = function() {
  return this.setRowsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.common.ReadData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupsList: jspb.Message.toObjectList(msg.getGroupsList(),
    proto.containersai.alameda.v1alpha1.datahub.common.Group.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.common.ReadData;
  return proto.containersai.alameda.v1alpha1.datahub.common.ReadData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.containersai.alameda.v1alpha1.datahub.common.Group;
      reader.readMessage(value,proto.containersai.alameda.v1alpha1.datahub.common.Group.deserializeBinaryFromReader);
      msg.addGroups(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.common.ReadData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.containersai.alameda.v1alpha1.datahub.common.Group.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Group groups = 1;
 * @return {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Group>}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.prototype.getGroupsList = function() {
  return /** @type{!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Group>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.containersai.alameda.v1alpha1.datahub.common.Group, 1));
};


/**
 * @param {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Group>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.prototype.setGroupsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Group=} opt_value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.prototype.addGroups = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.containersai.alameda.v1alpha1.datahub.common.Group, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ReadData} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.ReadData.prototype.clearGroupsList = function() {
  return this.setGroupsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.common.Row.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Row} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.toObject = function(includeInstance, msg) {
  var f, obj = {
    time: (f = msg.getTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    valuesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.common.Row;
  return proto.containersai.alameda.v1alpha1.datahub.common.Row.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Row} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTime(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addValues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.common.Row.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Row} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTime();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional google.protobuf.Timestamp time = 1;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.getTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 1));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.setTime = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.clearTime = function() {
  return this.setTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.hasTime = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated string values = 2;
 * @return {!Array<string>}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.getValuesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.setValuesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.addValues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Row.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.common.Group.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Group} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.toObject = function(includeInstance, msg) {
  var f, obj = {
    columnsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    rowsList: jspb.Message.toObjectList(msg.getRowsList(),
    proto.containersai.alameda.v1alpha1.datahub.common.Row.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.common.Group;
  return proto.containersai.alameda.v1alpha1.datahub.common.Group.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Group} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addColumns(value);
      break;
    case 2:
      var value = new proto.containersai.alameda.v1alpha1.datahub.common.Row;
      reader.readMessage(value,proto.containersai.alameda.v1alpha1.datahub.common.Row.deserializeBinaryFromReader);
      msg.addRows(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.common.Group.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Group} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getColumnsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getRowsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.containersai.alameda.v1alpha1.datahub.common.Row.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string columns = 1;
 * @return {!Array<string>}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.getColumnsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.setColumnsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.addColumns = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.clearColumnsList = function() {
  return this.setColumnsList([]);
};


/**
 * repeated Row rows = 2;
 * @return {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Row>}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.getRowsList = function() {
  return /** @type{!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Row>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.containersai.alameda.v1alpha1.datahub.common.Row, 2));
};


/**
 * @param {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Row>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.setRowsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Row=} opt_value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Row}
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.addRows = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.containersai.alameda.v1alpha1.datahub.common.Row, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.common.Group.prototype.clearRowsList = function() {
  return this.setRowsList([]);
};


goog.object.extend(exports, proto.containersai.alameda.v1alpha1.datahub.common);