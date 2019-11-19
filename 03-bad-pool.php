<?php

$client = new Memcached('bad_pool');
if ($client->isPristine()) {
    $client->setOption(Memcached::OPT_RETRY_TIMEOUT, 100);
    $client->addServer('fakememcached', 11211);
}

// the SET operation will hang forever on our fake memcached server
$client->set('foo', 'bar');
var_dump($client->getResultMessage());
