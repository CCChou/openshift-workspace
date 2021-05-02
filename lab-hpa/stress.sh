if [ $# -ne 2 ]; then
    echo "stress.sh [domain] [repeat_times]"
    exit 1
fi


DOMAIN=$1
COUNT=$2

for i in $(seq $COUNT);
do
    curl http://$DOMAIN/calc;
    echo ""
done

