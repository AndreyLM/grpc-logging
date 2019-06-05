<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: logging-service.proto

namespace V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * ---------------------- << USER LOGS >>
 *
 * Generated from protobuf message <code>v1.UserLog</code>
 */
class UserLog extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    private $id = 0;
    /**
     * Generated from protobuf field <code>int64 userId = 2;</code>
     */
    private $userId = 0;
    /**
     * Generated from protobuf field <code>int64 declarationId = 3;</code>
     */
    private $declarationId = 0;
    /**
     * Generated from protobuf field <code>string type = 4;</code>
     */
    private $type = '';
    /**
     * Generated from protobuf field <code>string message = 5;</code>
     */
    private $message = '';
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp createdAt = 6;</code>
     */
    private $createdAt = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *     @type int|string $userId
     *     @type int|string $declarationId
     *     @type string $type
     *     @type string $message
     *     @type \Google\Protobuf\Timestamp $createdAt
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\LoggingService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 userId = 2;</code>
     * @return int|string
     */
    public function getUserId()
    {
        return $this->userId;
    }

    /**
     * Generated from protobuf field <code>int64 userId = 2;</code>
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
     * Generated from protobuf field <code>int64 declarationId = 3;</code>
     * @return int|string
     */
    public function getDeclarationId()
    {
        return $this->declarationId;
    }

    /**
     * Generated from protobuf field <code>int64 declarationId = 3;</code>
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
     * Generated from protobuf field <code>string type = 4;</code>
     * @return string
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Generated from protobuf field <code>string type = 4;</code>
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
     * Generated from protobuf field <code>string message = 5;</code>
     * @return string
     */
    public function getMessage()
    {
        return $this->message;
    }

    /**
     * Generated from protobuf field <code>string message = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->message = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp createdAt = 6;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getCreatedAt()
    {
        return $this->createdAt;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp createdAt = 6;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->createdAt = $var;

        return $this;
    }

}

