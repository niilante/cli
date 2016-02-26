VAGRANTFILE_API_VERSION = '2'

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  # 構築に必要なリソース
  config.vm.provider 'virtualbox' do |v|
    v.memory = 2048
    v.cpus = 2
    v.name = 'arukas-cli'
    # 時刻同期
    v.customize ["setextradata", :id, "VBoxInternal/Devices/VMMDev/0/Config/GetHostTimeDisabled", 0]
  end



  # Ubuntu
  config.vm.box = "ubuntu/trusty64"

  # ここはNFSを使う。でなければrailsのパフォーマンスは大きく落ちる
  # Macではnfsを有効にすると、ファイルのオーナーまで共有されるので、
  # Vagrant上でdocker -v は使用できない場合がある
  config.vm.synced_folder '.', '/home/vagrant/src/github.com/arukasio/arukas-cli'

  config.vm.provision :shell, path: 'setup_vagrant.sh', privileged: false
end
