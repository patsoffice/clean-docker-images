# clean-docker-images

clean-docker-images is a CLI tool for identifying and removing unused, older docker images.

## Installing

### Directly From GitHub

Installing clean-docker-images is easy. First, use `go get` to install the latest version of the application
from GitHub. This command will install the `clean-docker-images` executable in the Go path:

    go get -u github.com/patsoffice/clean-docker-images

### Homebrew

    brew tap patsoffice/tools
    brew install clean-docker-images

## Running

Checking:

    $ clean-docker-images -e tcp://192.168.1.100:2375
    Can remove image sha256:c8743cd346eeea9f11ba10f472bccf7e066d6ccc1d3bd6c9c395847201de29f2 for mvance/unbound@sha256
    Can remove image sha256:f680154f76ee9c952bf68476f7b7f638d266d9527fb741fab7483f150ba3e723 for chronograf@sha256
    Can remove image sha256:1fc3f5bab93f707176e76dbf2020f61886f6549ed1214cad1ec7a1be4eb88f79 for linuxserver/letsencrypt@sha256
    Can remove image sha256:35de8cc24dfca8ad80fc93bba1bc7125240b3e27b457a1ef529aac5fb6217eef for influxdb@sha256
    Can remove image sha256:ce7e6cfbf54eb4cf66cf4491e045260a985372043a99ad2ef80fe2411c1a1502 for jgraph/drawio@sha256

Removing (add the `-r` or `--remove` flag):

    $ clean-docker-images -e tcp://192.168.1.100:2375 -r
    Can remove image sha256:c8743cd346eeea9f11ba10f472bccf7e066d6ccc1d3bd6c9c395847201de29f2 for mvance/unbound@sha256
    Removing image sha256:c8743cd346eeea9f11ba10f472bccf7e066d6ccc1d3bd6c9c395847201de29f2
    Can remove image sha256:f680154f76ee9c952bf68476f7b7f638d266d9527fb741fab7483f150ba3e723 for chronograf@sha256
    Removing image sha256:f680154f76ee9c952bf68476f7b7f638d266d9527fb741fab7483f150ba3e723
    Can remove image sha256:1fc3f5bab93f707176e76dbf2020f61886f6549ed1214cad1ec7a1be4eb88f79 for linuxserver/letsencrypt@sha256
    Removing image sha256:1fc3f5bab93f707176e76dbf2020f61886f6549ed1214cad1ec7a1be4eb88f79
    Can remove image sha256:35de8cc24dfca8ad80fc93bba1bc7125240b3e27b457a1ef529aac5fb6217eef for influxdb@sha256
    Removing image sha256:35de8cc24dfca8ad80fc93bba1bc7125240b3e27b457a1ef529aac5fb6217eef
    Can remove image sha256:ce7e6cfbf54eb4cf66cf4491e045260a985372043a99ad2ef80fe2411c1a1502 for jgraph/drawio@sha256
    Removing image sha256:ce7e6cfbf54eb4cf66cf4491e045260a985372043a99ad2ef80fe2411c1a1502
