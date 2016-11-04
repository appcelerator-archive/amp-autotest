docker service rm amp-test
docker service create --network amp-infra --name amp-autotest \
--restart-condition "none" \
--label io.amp.role="infrastructure" \
--mount type=bind,source=$GOPATH/src,target=/go/src \
appcelerator/amp-autotest:latest $1 $2
