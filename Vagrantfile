Vagrant.configure("2") do |config|
  config.vm.box = "debian/stretch64"
  config.vm.synced_folder ".", "/home/vagrant/llti"
  config.vm.provision "shell", privileged: true, inline: <<-SHELL
    # install Docker
    apt-get remove docker docker-engine docker.io

    apt-get update && apt-get install -y \
     apt-transport-https \
     ca-certificates \
     curl \
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

    apt-get update && apt-get install -y \
        docker-ce

    usermod -aG docker vagrant
  SHELL
end
