<?php
// GENERATED CODE -- DO NOT EDIT!

namespace V1;

/**
 * Services
 */
class UserLogServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \V1\CreateUserLogRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function CreateUserLog(\V1\CreateUserLogRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/v1.UserLogService/CreateUserLog',
        $argument,
        ['\V1\CreateUserLogResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \V1\ReadUserLogRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function ReadUserLog(\V1\ReadUserLogRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/v1.UserLogService/ReadUserLog',
        $argument,
        ['\V1\ReadUserLogResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \V1\FindUserLogsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function FindUserLogs(\V1\FindUserLogsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/v1.UserLogService/FindUserLogs',
        $argument,
        ['\V1\FindUserLogsResponse', 'decode'],
        $metadata, $options);
    }

}
