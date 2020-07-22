<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Parser;

/**
 */
class ParserIOClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Parser\Input $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Parser\Output
     */
    public function Parse(\Parser\Input $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/parser.ParserIO/Parse',
        $argument,
        ['\Parser\Output', 'decode'],
        $metadata, $options);
    }

}
