// source: datahub/applications/applications.proto
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

var datahub_common_queries_pb = require('../../datahub/common/queries_pb.js');
goog.object.extend(proto, datahub_common_queries_pb);
var datahub_common_rawdata_pb = require('../../datahub/common/rawdata_pb.js');
goog.object.extend(proto, datahub_common_rawdata_pb);
goog.exportSymbol('proto.prophetstor.datahub.applications.DeleteApplication', null, global);
goog.exportSymbol('proto.prophetstor.datahub.applications.ReadApplication', null, global);
goog.exportSymbol('proto.prophetstor.datahub.applications.WriteApplication', null, global);
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
proto.prophetstor.datahub.applications.WriteApplication = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.applications.WriteApplication, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.applications.WriteApplication.displayName = 'proto.prophetstor.datahub.applications.WriteApplication';
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
proto.prophetstor.datahub.applications.ReadApplication = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.datahub.applications.ReadApplication.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.datahub.applications.ReadApplication, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.applications.ReadApplication.displayName = 'proto.prophetstor.datahub.applications.ReadApplication';
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
proto.prophetstor.datahub.applications.DeleteApplication = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.datahub.applications.DeleteApplication.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.datahub.applications.DeleteApplication, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.applications.DeleteApplication.displayName = 'proto.prophetstor.datahub.applications.DeleteApplication';
}



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
proto.prophetstor.datahub.applications.WriteApplication.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.applications.WriteApplication.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.applications.WriteApplication} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.applications.WriteApplication.toObject = function(includeInstance, msg) {
  var f, obj = {
    measurement: jspb.Message.getFieldWithDefault(msg, 1, ""),
    writeData: (f = msg.getWriteData()) && datahub_common_rawdata_pb.WriteData.toObject(includeInstance, f)
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
 * @return {!proto.prophetstor.datahub.applications.WriteApplication}
 */
proto.prophetstor.datahub.applications.WriteApplication.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.applications.WriteApplication;
  return proto.prophetstor.datahub.applications.WriteApplication.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.applications.WriteApplication} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.applications.WriteApplication}
 */
proto.prophetstor.datahub.applications.WriteApplication.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMeasurement(value);
      break;
    case 2:
      var value = new datahub_common_rawdata_pb.WriteData;
      reader.readMessage(value,datahub_common_rawdata_pb.WriteData.deserializeBinaryFromReader);
      msg.setWriteData(value);
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
proto.prophetstor.datahub.applications.WriteApplication.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.applications.WriteApplication.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.applications.WriteApplication} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.applications.WriteApplication.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMeasurement();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getWriteData();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      datahub_common_rawdata_pb.WriteData.serializeBinaryToWriter
    );
  }
};


/**
 * optional string measurement = 1;
 * @return {string}
 */
proto.prophetstor.datahub.applications.WriteApplication.prototype.getMeasurement = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.applications.WriteApplication} returns this
 */
proto.prophetstor.datahub.applications.WriteApplication.prototype.setMeasurement = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional prophetstor.datahub.common.WriteData write_data = 2;
 * @return {?proto.prophetstor.datahub.common.WriteData}
 */
proto.prophetstor.datahub.applications.WriteApplication.prototype.getWriteData = function() {
  return /** @type{?proto.prophetstor.datahub.common.WriteData} */ (
    jspb.Message.getWrapperField(this, datahub_common_rawdata_pb.WriteData, 2));
};


/**
 * @param {?proto.prophetstor.datahub.common.WriteData|undefined} value
 * @return {!proto.prophetstor.datahub.applications.WriteApplication} returns this
*/
proto.prophetstor.datahub.applications.WriteApplication.prototype.setWriteData = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.datahub.applications.WriteApplication} returns this
 */
proto.prophetstor.datahub.applications.WriteApplication.prototype.clearWriteData = function() {
  return this.setWriteData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.datahub.applications.WriteApplication.prototype.hasWriteData = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.datahub.applications.ReadApplication.repeatedFields_ = [2];



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
proto.prophetstor.datahub.applications.ReadApplication.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.applications.ReadApplication.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.applications.ReadApplication} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.applications.ReadApplication.toObject = function(includeInstance, msg) {
  var f, obj = {
    measurement: jspb.Message.getFieldWithDefault(msg, 1, ""),
    whereConditionList: jspb.Message.toObjectList(msg.getWhereConditionList(),
    datahub_common_queries_pb.Condition.toObject, includeInstance)
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
 * @return {!proto.prophetstor.datahub.applications.ReadApplication}
 */
proto.prophetstor.datahub.applications.ReadApplication.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.applications.ReadApplication;
  return proto.prophetstor.datahub.applications.ReadApplication.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.applications.ReadApplication} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.applications.ReadApplication}
 */
proto.prophetstor.datahub.applications.ReadApplication.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMeasurement(value);
      break;
    case 2:
      var value = new datahub_common_queries_pb.Condition;
      reader.readMessage(value,datahub_common_queries_pb.Condition.deserializeBinaryFromReader);
      msg.addWhereCondition(value);
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
proto.prophetstor.datahub.applications.ReadApplication.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.applications.ReadApplication.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.applications.ReadApplication} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.applications.ReadApplication.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMeasurement();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getWhereConditionList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      datahub_common_queries_pb.Condition.serializeBinaryToWriter
    );
  }
};


/**
 * optional string measurement = 1;
 * @return {string}
 */
proto.prophetstor.datahub.applications.ReadApplication.prototype.getMeasurement = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.applications.ReadApplication} returns this
 */
proto.prophetstor.datahub.applications.ReadApplication.prototype.setMeasurement = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated prophetstor.datahub.common.Condition where_condition = 2;
 * @return {!Array<!proto.prophetstor.datahub.common.Condition>}
 */
proto.prophetstor.datahub.applications.ReadApplication.prototype.getWhereConditionList = function() {
  return /** @type{!Array<!proto.prophetstor.datahub.common.Condition>} */ (
    jspb.Message.getRepeatedWrapperField(this, datahub_common_queries_pb.Condition, 2));
};


/**
 * @param {!Array<!proto.prophetstor.datahub.common.Condition>} value
 * @return {!proto.prophetstor.datahub.applications.ReadApplication} returns this
*/
proto.prophetstor.datahub.applications.ReadApplication.prototype.setWhereConditionList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.datahub.common.Condition=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.datahub.common.Condition}
 */
proto.prophetstor.datahub.applications.ReadApplication.prototype.addWhereCondition = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.datahub.common.Condition, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.datahub.applications.ReadApplication} returns this
 */
proto.prophetstor.datahub.applications.ReadApplication.prototype.clearWhereConditionList = function() {
  return this.setWhereConditionList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.datahub.applications.DeleteApplication.repeatedFields_ = [2];



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
proto.prophetstor.datahub.applications.DeleteApplication.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.applications.DeleteApplication.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.applications.DeleteApplication} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.applications.DeleteApplication.toObject = function(includeInstance, msg) {
  var f, obj = {
    measurement: jspb.Message.getFieldWithDefault(msg, 1, ""),
    whereConditionList: jspb.Message.toObjectList(msg.getWhereConditionList(),
    datahub_common_queries_pb.Condition.toObject, includeInstance)
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
 * @return {!proto.prophetstor.datahub.applications.DeleteApplication}
 */
proto.prophetstor.datahub.applications.DeleteApplication.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.applications.DeleteApplication;
  return proto.prophetstor.datahub.applications.DeleteApplication.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.applications.DeleteApplication} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.applications.DeleteApplication}
 */
proto.prophetstor.datahub.applications.DeleteApplication.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMeasurement(value);
      break;
    case 2:
      var value = new datahub_common_queries_pb.Condition;
      reader.readMessage(value,datahub_common_queries_pb.Condition.deserializeBinaryFromReader);
      msg.addWhereCondition(value);
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
proto.prophetstor.datahub.applications.DeleteApplication.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.applications.DeleteApplication.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.applications.DeleteApplication} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.applications.DeleteApplication.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMeasurement();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getWhereConditionList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      datahub_common_queries_pb.Condition.serializeBinaryToWriter
    );
  }
};


/**
 * optional string measurement = 1;
 * @return {string}
 */
proto.prophetstor.datahub.applications.DeleteApplication.prototype.getMeasurement = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.prophetstor.datahub.applications.DeleteApplication} returns this
 */
proto.prophetstor.datahub.applications.DeleteApplication.prototype.setMeasurement = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated prophetstor.datahub.common.Condition where_condition = 2;
 * @return {!Array<!proto.prophetstor.datahub.common.Condition>}
 */
proto.prophetstor.datahub.applications.DeleteApplication.prototype.getWhereConditionList = function() {
  return /** @type{!Array<!proto.prophetstor.datahub.common.Condition>} */ (
    jspb.Message.getRepeatedWrapperField(this, datahub_common_queries_pb.Condition, 2));
};


/**
 * @param {!Array<!proto.prophetstor.datahub.common.Condition>} value
 * @return {!proto.prophetstor.datahub.applications.DeleteApplication} returns this
*/
proto.prophetstor.datahub.applications.DeleteApplication.prototype.setWhereConditionList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.datahub.common.Condition=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.datahub.common.Condition}
 */
proto.prophetstor.datahub.applications.DeleteApplication.prototype.addWhereCondition = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.datahub.common.Condition, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.datahub.applications.DeleteApplication} returns this
 */
proto.prophetstor.datahub.applications.DeleteApplication.prototype.clearWhereConditionList = function() {
  return this.setWhereConditionList([]);
};


goog.object.extend(exports, proto.prophetstor.datahub.applications);
