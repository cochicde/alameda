// source: datahub/rawdata/services.proto
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

var datahub_common_types_pb = require('../../datahub/common/types_pb.js');
goog.object.extend(proto, datahub_common_types_pb);
var datahub_rawdata_rawdata_pb = require('../../datahub/rawdata/rawdata_pb.js');
goog.object.extend(proto, datahub_rawdata_rawdata_pb);
var datahub_rawdata_types_pb = require('../../datahub/rawdata/types_pb.js');
goog.object.extend(proto, datahub_rawdata_types_pb);
var google_rpc_status_pb = require('../../google/rpc/status_pb.js');
goog.object.extend(proto, google_rpc_status_pb);
goog.exportSymbol('proto.prophetstor.datahub.rawdata.ReadRawdataRequest', null, global);
goog.exportSymbol('proto.prophetstor.datahub.rawdata.ReadRawdataResponse', null, global);
goog.exportSymbol('proto.prophetstor.datahub.rawdata.WriteRawdataRequest', null, global);
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
proto.prophetstor.datahub.rawdata.ReadRawdataRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.datahub.rawdata.ReadRawdataRequest.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.datahub.rawdata.ReadRawdataRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.rawdata.ReadRawdataRequest.displayName = 'proto.prophetstor.datahub.rawdata.ReadRawdataRequest';
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
proto.prophetstor.datahub.rawdata.ReadRawdataResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.datahub.rawdata.ReadRawdataResponse.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.datahub.rawdata.ReadRawdataResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.rawdata.ReadRawdataResponse.displayName = 'proto.prophetstor.datahub.rawdata.ReadRawdataResponse';
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
proto.prophetstor.datahub.rawdata.WriteRawdataRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.prophetstor.datahub.rawdata.WriteRawdataRequest.repeatedFields_, null);
};
goog.inherits(proto.prophetstor.datahub.rawdata.WriteRawdataRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.rawdata.WriteRawdataRequest.displayName = 'proto.prophetstor.datahub.rawdata.WriteRawdataRequest';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.repeatedFields_ = [2];



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
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.rawdata.ReadRawdataRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    databaseType: jspb.Message.getFieldWithDefault(msg, 1, 0),
    queriesList: jspb.Message.toObjectList(msg.getQueriesList(),
    datahub_rawdata_types_pb.Query.toObject, includeInstance)
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
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.rawdata.ReadRawdataRequest;
  return proto.prophetstor.datahub.rawdata.ReadRawdataRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.prophetstor.datahub.common.DatabaseType} */ (reader.readEnum());
      msg.setDatabaseType(value);
      break;
    case 2:
      var value = new datahub_rawdata_types_pb.Query;
      reader.readMessage(value,datahub_rawdata_types_pb.Query.deserializeBinaryFromReader);
      msg.addQueries(value);
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
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.rawdata.ReadRawdataRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabaseType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getQueriesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      datahub_rawdata_types_pb.Query.serializeBinaryToWriter
    );
  }
};


/**
 * optional prophetstor.datahub.common.DatabaseType database_type = 1;
 * @return {!proto.prophetstor.datahub.common.DatabaseType}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.getDatabaseType = function() {
  return /** @type {!proto.prophetstor.datahub.common.DatabaseType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.prophetstor.datahub.common.DatabaseType} value
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest} returns this
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.setDatabaseType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * repeated Query queries = 2;
 * @return {!Array<!proto.prophetstor.datahub.rawdata.Query>}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.getQueriesList = function() {
  return /** @type{!Array<!proto.prophetstor.datahub.rawdata.Query>} */ (
    jspb.Message.getRepeatedWrapperField(this, datahub_rawdata_types_pb.Query, 2));
};


/**
 * @param {!Array<!proto.prophetstor.datahub.rawdata.Query>} value
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest} returns this
*/
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.setQueriesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.datahub.rawdata.Query=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.datahub.rawdata.Query}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.addQueries = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.datahub.rawdata.Query, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataRequest} returns this
 */
proto.prophetstor.datahub.rawdata.ReadRawdataRequest.prototype.clearQueriesList = function() {
  return this.setQueriesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.repeatedFields_ = [2];



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
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.rawdata.ReadRawdataResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    status: (f = msg.getStatus()) && google_rpc_status_pb.Status.toObject(includeInstance, f),
    rawdataList: jspb.Message.toObjectList(msg.getRawdataList(),
    datahub_rawdata_rawdata_pb.ReadRawdata.toObject, includeInstance)
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
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.rawdata.ReadRawdataResponse;
  return proto.prophetstor.datahub.rawdata.ReadRawdataResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_rpc_status_pb.Status;
      reader.readMessage(value,google_rpc_status_pb.Status.deserializeBinaryFromReader);
      msg.setStatus(value);
      break;
    case 2:
      var value = new datahub_rawdata_rawdata_pb.ReadRawdata;
      reader.readMessage(value,datahub_rawdata_rawdata_pb.ReadRawdata.deserializeBinaryFromReader);
      msg.addRawdata(value);
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
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.rawdata.ReadRawdataResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStatus();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_rpc_status_pb.Status.serializeBinaryToWriter
    );
  }
  f = message.getRawdataList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      datahub_rawdata_rawdata_pb.ReadRawdata.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.rpc.Status status = 1;
 * @return {?proto.google.rpc.Status}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.getStatus = function() {
  return /** @type{?proto.google.rpc.Status} */ (
    jspb.Message.getWrapperField(this, google_rpc_status_pb.Status, 1));
};


/**
 * @param {?proto.google.rpc.Status|undefined} value
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} returns this
*/
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.setStatus = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} returns this
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.clearStatus = function() {
  return this.setStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.hasStatus = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated ReadRawdata rawdata = 2;
 * @return {!Array<!proto.prophetstor.datahub.rawdata.ReadRawdata>}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.getRawdataList = function() {
  return /** @type{!Array<!proto.prophetstor.datahub.rawdata.ReadRawdata>} */ (
    jspb.Message.getRepeatedWrapperField(this, datahub_rawdata_rawdata_pb.ReadRawdata, 2));
};


/**
 * @param {!Array<!proto.prophetstor.datahub.rawdata.ReadRawdata>} value
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} returns this
*/
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.setRawdataList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.datahub.rawdata.ReadRawdata=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdata}
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.addRawdata = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.datahub.rawdata.ReadRawdata, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.datahub.rawdata.ReadRawdataResponse} returns this
 */
proto.prophetstor.datahub.rawdata.ReadRawdataResponse.prototype.clearRawdataList = function() {
  return this.setRawdataList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.repeatedFields_ = [2];



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
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.rawdata.WriteRawdataRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    databaseType: jspb.Message.getFieldWithDefault(msg, 1, 0),
    rawdataList: jspb.Message.toObjectList(msg.getRawdataList(),
    datahub_rawdata_rawdata_pb.WriteRawdata.toObject, includeInstance)
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
 * @return {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest}
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.rawdata.WriteRawdataRequest;
  return proto.prophetstor.datahub.rawdata.WriteRawdataRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest}
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.prophetstor.datahub.common.DatabaseType} */ (reader.readEnum());
      msg.setDatabaseType(value);
      break;
    case 2:
      var value = new datahub_rawdata_rawdata_pb.WriteRawdata;
      reader.readMessage(value,datahub_rawdata_rawdata_pb.WriteRawdata.deserializeBinaryFromReader);
      msg.addRawdata(value);
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
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.rawdata.WriteRawdataRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabaseType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getRawdataList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      datahub_rawdata_rawdata_pb.WriteRawdata.serializeBinaryToWriter
    );
  }
};


/**
 * optional prophetstor.datahub.common.DatabaseType database_type = 1;
 * @return {!proto.prophetstor.datahub.common.DatabaseType}
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.getDatabaseType = function() {
  return /** @type {!proto.prophetstor.datahub.common.DatabaseType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.prophetstor.datahub.common.DatabaseType} value
 * @return {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest} returns this
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.setDatabaseType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * repeated WriteRawdata rawdata = 2;
 * @return {!Array<!proto.prophetstor.datahub.rawdata.WriteRawdata>}
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.getRawdataList = function() {
  return /** @type{!Array<!proto.prophetstor.datahub.rawdata.WriteRawdata>} */ (
    jspb.Message.getRepeatedWrapperField(this, datahub_rawdata_rawdata_pb.WriteRawdata, 2));
};


/**
 * @param {!Array<!proto.prophetstor.datahub.rawdata.WriteRawdata>} value
 * @return {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest} returns this
*/
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.setRawdataList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.prophetstor.datahub.rawdata.WriteRawdata=} opt_value
 * @param {number=} opt_index
 * @return {!proto.prophetstor.datahub.rawdata.WriteRawdata}
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.addRawdata = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.prophetstor.datahub.rawdata.WriteRawdata, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.prophetstor.datahub.rawdata.WriteRawdataRequest} returns this
 */
proto.prophetstor.datahub.rawdata.WriteRawdataRequest.prototype.clearRawdataList = function() {
  return this.setRawdataList([]);
};


goog.object.extend(exports, proto.prophetstor.datahub.rawdata);