
1. Pull ubuntu image
```
docker pull ubuntu
```
2. Create ubuntu container by name ansible-master
```
docker run --name ansible-master  -it ubuntu
```
3. Create ubuntu container by name ansible-node
```
docker run -it --name ansible-node ubuntu
```
4. Update ansible-master  & ansible-node containers 
```
apt-get update
```
5. Install python3 on ansible-master & ansible-node containers
```
apt-get install python3
```
 6. Install ansible on ansible-master
```
 apt-get install ansible
```
7. Check ansible is installed
```
ansible --version
```
8. Create ansible config 
```
ansible-config init --disable >ansible.cfg
```
9. Install nano on ansible-master and ansible-node
```
apt-get install nano
```
10. Edit ansible.cfg on ansible-master to add defaults inventory = /etc/ansible/hosts.int
```
nano ./ansible.cfg
```

```
[defaults]
inventory = /etc/ansible/hosts.ini
```
10. Update hosts.ini file with ansible-node container ip address, to get IP of ansible-node use below command
```
hostname -I
```

```
mkdir /etc/ansible
```

```
touch /etc/ansible/hosts.ini
```

```
nano /etc/ansible/hosts.ini
```

```
[Group1]
172.17.0.3
```
11. save the files and exit
```
ctrl + o + enter
```

```
ctrl + x
```
12. Generate ssh keys on ansible-master & ansible-node without passpharse
```
ssh-keygen -t -rsa
```

```
apt-get install openssh-server
```

```
ssh-keygen
```
13. SSH copy id
```
ssh-copy-id -i ~/.ssh/id_ed25519.pub root@172.17.0.03
```
14. Above command fails follow below steps
on ansible-master
```
cat ~/.ssh/id_ed25519.pub
```

```
copy the public key
```
on asible-node

```
touch ~/.ssh/authorized_keys
```

```
nano ~/.ssh/authorized_keys
```
paste the ansible-master public key in authorized_key ctrl + o, enter ctrl + x 
15. check the ssh service is working 
```
service ssh status 
```

```
service ssh start
```
17. Check if password less connection is established or not 
```
ansible 172.17.0.3 -m ping
```

```
ssh-copy-id -f "-o IdentityFile ec2.pem" ubuntu@3.106.201.11
```
