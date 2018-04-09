Vagrant.configure("2") do |config|
  config.vm.box = "debian/stretch64"

  config.vm.synced_folder ".", "/home/vagrant/go/src/github.com/d-led/linux_limits_to_influxdb"

  config.vm.provision "shell", privileged: true, inline: <<-SHELL
    # install Docker
    apt-get remove docker docker-engine docker.io

    apt-get update && apt-get install -y \
     apt-transport-https \
     ca-certificates \
     curl \
     wget \
     git \
     gnupg2 \
     software-properties-common

    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
    apt-key fingerprint 0EBFCD88
    add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/debian \
       $(lsb_release -cs) \
       stable"

    curl -L https://github.com/docker/compose/releases/download/1.20.1/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose

    apt-get update && apt-get install -y \
        docker-ce

    usermod -aG docker vagrant

    # install Golang
    mkdir tmp
    cd tmp
    wget -qq https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz
    tar -xvf go1.10.1.linux-amd64.tar.gz
    mv go /usr/local
    cd ..

    # profile defaults
    echo export GOROOT=/usr/local/go >> /home/vagrant/.profile
    echo export GOPATH=/home/vagrant/go >> /home/vagrant/.profile
    echo cd /home/vagrant/go/src/github.com/d-led/linux_limits_to_influxdb >> /home/vagrant/.profile
    echo export PATH=\\\$GOPATH/bin:\\\$GOROOT/bin:\\\$PATH >> /home/vagrant/.profile

    chown -R vagrant /home/vagrant
  SHELL
end
