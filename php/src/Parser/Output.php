<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/prase.proto

namespace Parser;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>parser.Output</code>
 */
class Output extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string Title = 1;</code>
     */
    private $Title;
    /**
     * Generated from protobuf field <code>repeated string Canonical = 2;</code>
     */
    private $Canonical;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $Title
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $Canonical
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Prase::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string Title = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getTitle()
    {
        return $this->Title;
    }

    /**
     * Generated from protobuf field <code>repeated string Title = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setTitle($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->Title = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string Canonical = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCanonical()
    {
        return $this->Canonical;
    }

    /**
     * Generated from protobuf field <code>repeated string Canonical = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCanonical($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->Canonical = $arr;

        return $this;
    }

}

