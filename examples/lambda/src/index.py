import os
from pymemcache.client.base import Client

MEMCACHED_CLUSTER_ADDRESS = os.getenv('MEMCACHED_CLUSTER_ADDRESS')

mc = Client(MEMCACHED_CLUSTER_ADDRESS)


def handler(event, context):
    mc.set('foo', 'bar')
    result = mc.get('foo')
    assert result == b'bar'
