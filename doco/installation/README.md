## Install via Ansible
#### NOTE the manual setup steps
1. You need a working local GOLang environment for ansible to build the project
2. Update ansible_deploy.yml with your local gopath
   ```
    gopath: "/Users/rebelit/go"
    goroot: "/usr/local/opt/go/libexec"
   ```
3. Update the `ansible_host` file with your DNS name or IP
4. Run it (depending on your local setup you may not need `--ask-sudo-pass`)
    ```
    ansible-playbook ansible_deploy.yml -i ansible_host --ask-sudo-pass
    ```

---

## Build with Docker
1. Build the image then run docker to compile the app binary.  no need for GOLang to be configured
locally.
    ```
    ensure you are in the repo directory
    docker build -t rpiot .
    docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t --rm rpiot /build.sh
    scp main pi@yourHost:/your/directory/here/
    ```
2. run it however you want `. /main`
3. create a service file manually and daemonize it

---

## Build with make
1. process yet to be created
