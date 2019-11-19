<?php

$client = new Memcached();
$client->addServer('fakememcached', 11211);

// the SET operation will hang forever on our fake memcached server
$client->set('foo', 'bar');
var_dump($client->getResultMessage());
