There are two methods one is through ssh other is password  
  
1. Making ansible-node1 password less using SSH  
1. Create ansible-node1,ansible-node2 in AWS account  
3. Attach keypair for both the instances  
3. downloads ec2.pem file  
4. chmod 600 ~/downloads/ec2.pem  
```  
ssh-copy-id -f "-o IdentityFile ~/Developer/Mallikarjun/ec2.pem" ubuntu@65.2.70.183
```  

```
ssh-copy-id -f "-o IdentityFile ~/Developer/Mallikarjun/ec2.pem" ec2-user@13.127.72.186
```

```
ssh-copy-id -f "-o IdentityFile ~/Developer/Mallikarjun/ec2.pem" admin@3.110.154.212
```

```  
ssh ubuntu@13.235.0.230
```  
  
\  
2. Making ansible-node2 password less using password  
1. connect to second node  
```  
ssh -i ~/downloads/ec2.pem ubuntu@13.201.92.133  
```  
  
2. Go to the file   
```  
sudo vim  /etc/ssh/sshd_config.d/60-cloudimg-settings.conf  
```  
  
3.  Update PasswordAuthentication yes  
4.  uncommanet in sshd_config PasswordAuthentication Yes  
```  
sudo vim /etc/ssh/sshd_config  
```  
6. Uncomment  PasswordAuthentication yes  
7. Set password  
```  
sudo passwd ubuntu  
```  
8. Restart SSH   
```  
sudo systemctl restart ssh  
```  
9. exit  
10. do ssh-copy id    
```  
shh-copy-id [ubuntu@13.201.92.133](mailto:ubuntu@13.201.92.133)  
  
```  
11. Enter password  
  
```  
touch inventory.ini  
```  
  
```  
vim inventory.ini  
```  
  
12. Add user@ipaddress  
  
  
```  
[ubuntu@13.232.5.6](mailto:ubuntu@13.232.5.6)  
[ubuntu@35.154.15.12](mailto:ubuntu@35.154.15.12)  
```  
13. Save file  
14. Test connection  
```  
ansible -i hosts.ini -m ping all  
```  
  
14. Ansible adhoc commands  
  
```  
ansible -i hosts.ini  -m shell -a "sudo apt-get update" all  
```  
  
  
```  
 ansible -i hosts.ini  -m shell -a "sudo ls ~/.ssh" all  
```

In Inventory file

ipaddress ansible_ssh_private_key_file=~/.ssh/demo

```
sudo ps -ef | grep apache2
```

```
sudo systemctl status apache2
```