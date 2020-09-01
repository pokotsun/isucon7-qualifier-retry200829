branch=master
build: 
        (cd /home/isucon/isubata && git reset --hard HEAD && git fetch && git checkout ${branch} && git merge origin ${branch})
        (cd /home/isucon/isubata/webapp/go && make)
        sudo systemctl restart nginx
        sudo systemctl restart isubata.golang
analyze:
        sudo cp /dev/null /var/log/mysql/mysql-slow.log
        sudo cp /dev/null /var/log/nginx/access.log
        (cd bench && ./bin/bench -remotes=127.0.0.1 -output result.json)
        sudo alp --file=/var/log/nginx/access.log ltsv -r --sort sum | head -n 30
.PHONY: build analyze
