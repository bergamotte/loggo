echo "Spawning 20 processes"
for i in {1..20} ;
do
    ( ./loggo -pid="./tmp/loggo$i.pid" & )
done
