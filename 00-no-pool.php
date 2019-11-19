<?php

$client = new Memcached();
$client->addServer('fakememcached', 11211);

var_dump($client->get('foo'));
