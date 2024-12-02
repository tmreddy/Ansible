In Ansible, **roles** and **tasks** are two core concepts that help organize, modularize, and structure playbooks efficiently. Understanding these concepts is crucial to writing clean, reusable, and maintainable Ansible code.

### 1. **Ansible Tasks**
A **task** is the smallest unit of work in Ansible. Each task performs a single operation, such as installing a package, starting a service, or copying a file.

#### Key Features of Tasks:
- A task is defined in the `tasks:` section of a playbook or a role.
- Tasks are executed sequentially within a playbook.
- Tasks can use Ansible modules to perform specific operations (e.g., `ansible.builtin.yum`, `ansible.builtin.service`, `ansible.builtin.file`).
- Each task can include a `name` to describe what it is doing.

#### Task Syntax:
```yaml
- name: Install Nginx
  ansible.builtin.yum:
    name: nginx
    state: present
```

#### Example: Basic Task in a Playbook
```yaml
---
- name: Install and configure nginx
  hosts: web_servers
  become: yes  # Run as sudo
  tasks:
    - name: Install Nginx package
      ansible.builtin.yum:
        name: nginx
        state: present

    - name: Start Nginx service
      ansible.builtin.service:
        name: nginx
        state: started

    - name: Enable Nginx service on boot
      ansible.builtin.systemd:
        name: nginx
        enabled: yes
```

In this example:
- The first task installs the Nginx package using the `yum` module.
- The second task ensures that the Nginx service is started using the `service` module.
- The third task ensures that Nginx is enabled to start on boot using the `systemd` module.

### 2. **Ansible Roles**
A **role** in Ansible is a way of organizing tasks, variables, templates, files, and handlers into a structured and reusable format. Roles allow you to split complex playbooks into reusable components, making them easier to maintain and share.

Roles are designed to:
- Organize tasks into a structured directory layout.
- Allow for the reuse of common functionality.
- Support automatic variable loading, templates, and handlers.

Roles can be used to encapsulate tasks that configure specific aspects of a system, such as setting up a web server, database server, or application.

#### Role Directory Structure:
A typical role follows a predefined directory structure:

```
roles/
  └── my_role/
      ├── tasks/
      │   └── main.yml
      ├── handlers/
      │   └── main.yml
      ├── defaults/
      │   └── main.yml
      ├── vars/
      │   └── main.yml
      ├── templates/
      ├── files/
      └── meta/
          └── main.yml
```

##### Key Components of a Role:
- **tasks/**: Contains the tasks that the role will execute. The main entry point for tasks is usually `main.yml`.
- **handlers/**: Contains handlers (special tasks that are triggered by notifications).
- **defaults/**: Contains default variables for the role.
- **vars/**: Contains custom variables for the role (can override defaults).
- **templates/**: Contains Jinja2 templates to be used in the role.
- **files/**: Contains static files to be copied to managed nodes.
- **meta/**: Contains metadata about the role (dependencies, author, etc.).

#### Example: Basic Role Structure for Installing Nginx
Suppose you want to create a role for installing and configuring Nginx. The directory structure might look like this:

```
roles/
  └── nginx/
      ├── tasks/
      │   └── main.yml
      ├── handlers/
      │   └── main.yml
      ├── defaults/
      │   └── main.yml
      ├── vars/
      │   └── main.yml
```

##### **tasks/main.yml**:
```yaml
---
# tasks file for nginx

- name: Install Nginx
  ansible.builtin.yum:
    name: nginx
    state: present

- name: Start and enable Nginx service
  ansible.builtin.systemd:
    name: nginx
    state: started
    enabled: yes
```

##### **defaults/main.yml**:
```yaml
---
# Default variables for nginx role
nginx_package_name: nginx
```

##### **handlers/main.yml**:
```yaml
---
# Handlers file for nginx

- name: restart nginx
  ansible.builtin.systemd:
    name: nginx
    state: restarted
```

##### **vars/main.yml** (optional):
```yaml
---
# Custom variables for nginx
nginx_port: 80
```

#### Example Playbook Using the Role:
Once the role is created, you can use it in a playbook:

```yaml
---
- name: Configure web servers
  hosts: web_servers
  become: yes
  roles:
    - nginx
```

In this example, the `nginx` role will be applied to all hosts in the `web_servers` group. Ansible will automatically look for a `nginx/tasks/main.yml` file and execute the tasks defined there.

### 3. **Tasks within Roles**
In a role, tasks are organized into separate files, but typically the `main.yml` file serves as the main entry point. The tasks in a role can include:
- Package installations.
- Configuration file modifications.
- Service management (start/stop/restart).
- File copying or template rendering.

You can include additional task files using `include_tasks` or `import_tasks` to modularize the tasks further.

#### Example: Including Additional Task Files in a Role
```yaml
---
# tasks/main.yml
- name: Install Nginx
  ansible.builtin.yum:
    name: nginx
    state: present

- name: Include the security tasks
  include_tasks: security.yml
```

In this example, after installing Nginx, the role includes additional tasks from the `security.yml` file, which could contain tasks like configuring a firewall, applying security patches, etc.

### 4. **Using Variables with Roles**
Roles support variable files like `defaults/main.yml` and `vars/main.yml`. You can define default values in `defaults/main.yml` and more specific variables in `vars/main.yml`.

#### Example: Role Variables in Action
Let's use the `nginx` role and define a variable for the Nginx port:

```yaml
# vars/main.yml
nginx_port: 8080
```

```yaml
# tasks/main.yml
- name: Install Nginx
  ansible.builtin.yum:
    name: nginx
    state: present

- name: Configure Nginx port
  ansible.builtin.template:
    src: nginx.conf.j2
    dest: /etc/nginx/nginx.conf
```

In the template file (`nginx.conf.j2`), you can use the variable `nginx_port`:

```nginx
server {
    listen {{ nginx_port }};
    server_name localhost;

    location / {
        root /usr/share/nginx/html;
        index index.html index.htm;
    }
}
```

When the playbook runs, Ansible will substitute the value of `nginx_port` into the template, rendering the configuration with the correct port.

### 5. **Handlers in Roles**
Handlers are special tasks that only run when notified by another task. They are typically used to perform actions like restarting a service after a configuration change.

#### Example: Handler for Restarting Nginx
In the `handlers/main.yml` of the `nginx` role:

```yaml
---
# handlers/main.yml
- name: restart nginx
  ansible.builtin.systemd:
    name: nginx
    state: restarted
```

Then, in the tasks section, you can notify the handler when a task changes something (like a config file):

```yaml
- name: Configure Nginx
  ansible.builtin.template:
    src: nginx.conf.j2
    dest: /etc/nginx/nginx.conf
  notify:
    - restart nginx
```

This ensures that the Nginx service is restarted only when the configuration file has changed.

### 6. **Role Dependencies**
Roles can depend on other roles, which makes them even more reusable and composable. You can declare role dependencies in the `meta/main.yml` file.

#### Example: Role Dependencies in `meta/main.yml`
```yaml
---
dependencies:
  - role: nginx
  - role: firewall
```

In this example, before executing the tasks in the current role, Ansible will ensure that the `nginx` and `firewall` roles are executed first.

### Conclusion
- **Tasks**: The basic building blocks in Ansible, representing a single unit of work (e.g., installing a package, managing a service). Tasks are organized under the `tasks:` section of a playbook or role.
- **Roles**: A way to organize tasks and related files (templates, variables, etc.) into reusable and modular units. Roles help structure your playbooks, making them easier to maintain and share.
  
By using roles and tasks effectively, you can write more modular, reusable, and maintainable Ansible automation that is easy to scale and extend.