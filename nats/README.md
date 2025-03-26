[NATS Docs](https://docs.nats.io/)

helm install --values values.yaml nats nats/nats

pub -s nats://localhost:4222 idol.wm "hello"

nats request idol.wm "Hello I am requesting some messages"
nats reply idol.wm "I am client 1"

nats pub idol.wm "hello"
nats sub idol.wm

nats account info
nats stream add my_stream
nats stream info my_stream

nats pub foo --count=1000 --sleep 1s "publication #{{Count}} @ {{TimeStamp}}"

nats consumer add
nats consumer next my_stream my_consumer --count 1000 --timeout=5m

nats % nats kv add my-kv
nats % nats kv put my-kv k1 v1alpha
nats % nats kv ls my-kv            
nats % nats kv get my-kv k1        
nats % nats kv del my-kv k1        
