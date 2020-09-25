// source: datahub/licenses/services.proto
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

var datahub_licenses_licenses_pb = require('../../datahub/licenses/licenses_pb.js');
goog.object.extend(proto, datahub_licenses_licenses_pb);
var google_rpc_status_pb = require('../../google/rpc/status_pb.js');
goog.object.extend(proto, google_rpc_status_pb);
goog.exportSymbol('proto.prophetstor.datahub.licenses.GetLicenseResponse', null, global);
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
proto.prophetstor.datahub.licenses.GetLicenseResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.prophetstor.datahub.licenses.GetLicenseResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.prophetstor.datahub.licenses.GetLicenseResponse.displayName = 'proto.prophetstor.datahub.licenses.GetLicenseResponse';
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
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.prophetstor.datahub.licenses.GetLicenseResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.prophetstor.datahub.licenses.GetLicenseResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    status: (f = msg.getStatus()) && google_rpc_status_pb.Status.toObject(includeInstance, f),
    license: (f = msg.getLicense()) && datahub_licenses_licenses_pb.License.toObject(includeInstance, f)
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
 * @return {!proto.prophetstor.datahub.licenses.GetLicenseResponse}
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.prophetstor.datahub.licenses.GetLicenseResponse;
  return proto.prophetstor.datahub.licenses.GetLicenseResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.prophetstor.datahub.licenses.GetLicenseResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.prophetstor.datahub.licenses.GetLicenseResponse}
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = new datahub_licenses_licenses_pb.License;
      reader.readMessage(value,datahub_licenses_licenses_pb.License.deserializeBinaryFromReader);
      msg.setLicense(value);
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
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.prophetstor.datahub.licenses.GetLicenseResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.prophetstor.datahub.licenses.GetLicenseResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStatus();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_rpc_status_pb.Status.serializeBinaryToWriter
    );
  }
  f = message.getLicense();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      datahub_licenses_licenses_pb.License.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.rpc.Status status = 1;
 * @return {?proto.google.rpc.Status}
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.getStatus = function() {
  return /** @type{?proto.google.rpc.Status} */ (
    jspb.Message.getWrapperField(this, google_rpc_status_pb.Status, 1));
};


/**
 * @param {?proto.google.rpc.Status|undefined} value
 * @return {!proto.prophetstor.datahub.licenses.GetLicenseResponse} returns this
*/
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.setStatus = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.datahub.licenses.GetLicenseResponse} returns this
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.clearStatus = function() {
  return this.setStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.hasStatus = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional License license = 2;
 * @return {?proto.prophetstor.datahub.licenses.License}
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.getLicense = function() {
  return /** @type{?proto.prophetstor.datahub.licenses.License} */ (
    jspb.Message.getWrapperField(this, datahub_licenses_licenses_pb.License, 2));
};


/**
 * @param {?proto.prophetstor.datahub.licenses.License|undefined} value
 * @return {!proto.prophetstor.datahub.licenses.GetLicenseResponse} returns this
*/
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.setLicense = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.prophetstor.datahub.licenses.GetLicenseResponse} returns this
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.clearLicense = function() {
  return this.setLicense(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.prophetstor.datahub.licenses.GetLicenseResponse.prototype.hasLicense = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.prophetstor.datahub.licenses);
