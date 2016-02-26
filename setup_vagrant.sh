sudo apt-get update
sudo apt-get install -y git
curl -O https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
tar zxf go1.5.2.linux-amd64.tar.gz
echo 'export PATH="$PATH:$HOME/go/bin:$HOME/bin"' >> /home/vagrant/.profile
echo 'export GOROOT="$HOME/go"' >> /home/vagrant/.profile
echo 'export GOPATH="$HOME"' >> /home/vagrant/.profile
echo 'export GO15VENDOREXPERIMENT=1' >> /home/vagrant/.profile
exec $SHELL -l
sudo chown -R vagrant:vagrant $HOME
go get github.com/kr/godep

