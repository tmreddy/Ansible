### Ansible Template – In Detail

Ansible is an automation tool used for configuration management, application deployment, task automation, and multi-node orchestration. One of the most powerful features of Ansible is its ability to **render dynamic configurations** using templates. 

Templates in Ansible are based on **Jinja2** templating, which allows you to embed Python-like expressions and logic into files (typically configuration files) to be rendered on remote machines.

### What is an Ansible Template?

An Ansible **template** is a file (usually a configuration file) that contains placeholders for variables or conditional logic. These placeholders are dynamically replaced with values when the template is applied. 

Ansible templates are typically stored as **.j2** files (denoting Jinja2 format). These templates can contain:
- **Variables**: Data passed into the template from Ansible variables.
- **Loops**: Repeating patterns based on a list or dictionary.
- **Conditionals**: Logic to handle "if-else" situations.

Ansible’s **`template`** module is used to copy and render a template file from the local control machine to the remote target.

### Common Use Cases for Ansible Templates

1. **Dynamic Configuration Files**:
   - Generate configuration files that are personalized based on variables (e.g., server names, IP addresses, environment settings).
   
2. **Generating Secrets or Tokens**:
   - Create files that include sensitive data (like API tokens) injected via variables.
   
3. **Scripting & Automation**:
   - Generate scripts for automation, such as start-up scripts, monitoring configurations, etc.

4. **Environment-Specific Configurations**:
   - Use templates to define environment-specific configurations, such as dev, staging, or prod.

### Syntax and How It Works

#### 1. **Jinja2 Templating Syntax**
Jinja2 templates use **`{{ }}`** for variable substitution, **`{% %}`** for control structures (like loops or conditionals), and **`{# #}`** for comments.

- **Variables**: `{{ variable_name }}`
- **Conditionals**: `{% if condition %} ... {% endif %}`
- **Loops**: `{% for item in items %} ... {% endfor %}`

#### 2. **Ansible Template Module**

Ansible uses the `template` module to apply Jinja2 templates to files.

Here’s an example playbook that uses the `template` module:

```yaml
---
- name: Example of Ansible Template
  hosts: webservers
  vars:
    app_name: myapp
    app_version: 1.2.3
    server_name: "{{ ansible_hostname }}"
  tasks:
    - name: Deploy configuration file from template
      template:
        src: /path/to/template.conf.j2
        dest: /etc/myapp/config.conf
```

Here, the `src` is the local template file (in Jinja2 format), and `dest` is the path where the rendered file will be placed on the target machine.

### Detailed Example: Web Server Configuration

Let's walk through an example where we want to deploy a configuration file to a web server. We'll use a template to render the web server configuration dynamically based on variables.

#### Step 1: Define the Template

Create a Jinja2 template for a web server configuration file (`webserver.conf.j2`):

**`webserver.conf.j2`**:

```jinja
# Configuration for {{ app_name }} - version {{ app_version }}
server {
    listen 80;
    server_name {{ server_name }};
    
    location / {
        proxy_pass http://127.0.0.1:8080;
    }
    
    {% if app_version == '1.2.3' %}
    # Special configuration for version 1.2.3
    root /var/www/{{ app_name }}/v1.2.3;
    {% else %}
    root /var/www/{{ app_name }}/latest;
    {% endif %}
}
```

In this template:
- `{{ app_name }}` and `{{ app_version }}` will be replaced by the corresponding variables defined in the Ansible playbook.
- The `server_name` is dynamically set to the host’s `ansible_hostname`.
- The `if-else` logic is used to customize the root directory based on the application version.

#### Step 2: Create the Ansible Playbook

Create an Ansible playbook to render this template and deploy the file to the web server.

**`deploy-webserver.yml`**:

```yaml
---
- name: Deploy Web Server Configuration
  hosts: webservers
  vars:
    app_name: "myapp"
    app_version: "1.2.3"
    server_name: "{{ ansible_hostname }}"
  tasks:
    - name: Deploy the web server configuration
      template:
        src: webserver.conf.j2
        dest: /etc/nginx/sites-available/myapp.conf
      notify:
        - Restart Nginx

  handlers:
    - name: Restart Nginx
      service:
        name: nginx
        state: restarted
```

#### Explanation of the Playbook:
- The **`vars`** section defines variables that are passed to the template (such as `app_name`, `app_version`, and `server_name`).
- The **`template`** module renders the `webserver.conf.j2` file and copies it to the target machine's `/etc/nginx/sites-available/myapp.conf`.
- The **`notify`** directive triggers the handler to restart Nginx if the template is applied successfully.

#### Step 3: Apply the Playbook

Run the playbook with the following command:

```bash
ansible-playbook deploy-webserver.yml
```

### Advanced Features of Ansible Templates

#### 1. **Looping in Templates**

You can loop over lists or dictionaries and dynamically create entries in configuration files.

**Example:**
Let's say you want to generate a list of servers in a configuration file.

**`servers.conf.j2`**:

```jinja
# List of servers
servers:
  {% for server in server_list %}
  - {{ server.name }}: {{ server.ip }}
  {% endfor %}
```

**Playbook**:

```yaml
---
- name: Generate server list
  hosts: localhost
  vars:
    server_list:
      - { name: "server1", ip: "192.168.1.1" }
      - { name: "server2", ip: "192.168.1.2" }
  tasks:
    - name: Generate servers configuration
      template:
        src: servers.conf.j2
        dest: /tmp/servers.conf
```

This playbook will loop over `server_list` and render a list of servers in the `servers.conf` file.

#### 2. **Using Ansible Facts**

You can use **Ansible facts** (collected system information) in your templates. For example, the `ansible_hostname` fact can be used to dynamically insert the hostname of the machine.

Example template:

```jinja
# System info for {{ ansible_hostname }}
hostname = {{ ansible_hostname }}
ip_address = {{ ansible_default_ipv4.address }}
```

### Conclusion

Ansible’s **template** module, combined with Jinja2 templating, offers a flexible and powerful way to dynamically generate configuration files based on variable data. By leveraging **loops**, **conditionals**, and **Ansible facts**, you can create highly customizable configurations that adapt to different environments and use cases.

Templates help ensure consistency and automation while simplifying configuration management for large-scale infrastructure.