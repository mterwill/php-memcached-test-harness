<?php

$client = new Memcached('good_pool');
if ($client->isPristine()) {
    $client->addServer('fakememcached', 11211);
}

var_dump($client->get('foo'));
