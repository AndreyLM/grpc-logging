<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: logging-service.proto

namespace V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * FIND LOGS
 *
 * Generated from protobuf message <code>v1.FindUserLogsRequest</code>
 */
class FindUserLogsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string api = 1;</code>
     */
    private $api = '';
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp from = 2;</code>
     */
    private $from = null;
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp to = 3;</code>
     */
    private $to = null;
    /**
     * Generated from protobuf field <code>int64 userId = 4;</code>
     */
    private $userId = 0;
    /**
     * Generated from protobuf field <code>int64 declarationId = 5;</code>
     */
    private $declarationId = 0;
    /**
     * Generated from protobuf field <code>string type = 6;</code>
     */
    private $type = '';
    /**
     * Generated from protobuf field <code>int32 limit = 7;</code>
     */
    private $limit = 0;
    /**
     * Generated from protobuf field <code>int32 offset = 8;</code>
     */
    private $offset = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $api
     *     @type \Google\Protobuf\Timestamp $from
     *     @type \Google\Protobuf\Timestamp $to
     *     @type int|string $userId
     *     @type int|string $declarationId
     *     @type string $type
     *     @type int $limit
     *     @type int $offset
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\LoggingService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string api = 1;</code>
     * @return string
     */
    public function getApi()
    {
        return $this->api;
    }

    /**
     * Generated from protobuf field <code>string api = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setApi($var)
    {
        GPBUtil::checkString($var, True);
        $this->api = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp from = 2;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getFrom()
    {
        return $this->from;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp from = 2;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setFrom($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->from = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp to = 3;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getTo()
    {
        return $this->to;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp to = 3;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setTo($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->to = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 userId = 4;</code>
     * @return int|string
     */
    public function getUserId()
    {
        return $this->userId;
    }

    /**
     * Generated from protobuf field <code>int64 userId = 4;</code>
     * @param int|string $var
     * @return $this
     */
    public function setUserId($var)
    {
        GPBUtil::checkInt64($var);
        $this->userId = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 declarationId = 5;</code>
     * @return int|string
     */
    public function getDeclarationId()
    {
        return $this->declarationId;
    }

    /**
     * Generated from protobuf field <code>int64 declarationId = 5;</code>
     * @param int|string $var
     * @return $this
     */
    public function setDeclarationId($var)
    {
        GPBUtil::checkInt64($var);
        $this->declarationId = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string type = 6;</code>
     * @return string
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Generated from protobuf field <code>string type = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setType($var)
    {
        GPBUtil::checkString($var, True);
        $this->type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 7;</code>
     * @return int
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 7;</code>
     * @param int $var
     * @return $this
     */
    public function setLimit($var)
    {
        GPBUtil::checkInt32($var);
        $this->limit = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 8;</code>
     * @return int
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 8;</code>
     * @param int $var
     * @return $this
     */
    public function setOffset($var)
    {
        GPBUtil::checkInt32($var);
        $this->offset = $var;

        return $this;
    }

}

