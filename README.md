# hasselhoffme
```bash
wget http://tiny.cc/hasselhoff -O - | bash
```
```bash
curl -L http://tiny.cc/hasselhoff | bash
```
## Requirements:

- debian/ubuntu:

```bash
sudo apt-get update && sudo apt-get install -y libxml2-dev
```

- fedora/redhat:

```bash
sudo yum install libxml2-devel
```

For releases in Github you have to set an environment variable in travis: ```GITHUB_RELEASES_API_KEY``` which its value
is set running the ```travis``` gem. To install and execute it just follow the next steps:

- In debian/ubuntu

```bash
sudo apt-get update && sudo apt-get install ruby-dev
```

- In fedora/redhat

```bash
sudo yum install ruby-devel rubygems
```

```bash
sudo gem install travis
```

Ensure you have it installed running ```travis version```

```bash
travis setup releases --force
```

And then it will create an api key for deploying your app in GitHub