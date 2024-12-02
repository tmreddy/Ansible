### Creating an Ansible Role to Install Apache2 on Ubuntu

In this example, we will create an Ansible role to install and configure **Apache2** on an **Ubuntu** system. We will use the standard Ansible role structure, which includes folders like `defaults`, `files`, `handlers`, `meta`, `tasks`, `templates`, `vars`, and `tests`. Each folder will serve a specific purpose, and we will explain its contents with examples.

### Step-by-Step Guide to Create the Role

#### 1. **Create Role Directory Structure**

First, we need to create the directory structure for the role.

```bash
$ mkdir -p roles/apache2/{defaults,files,handlers,meta,tasks,templates,vars,tests}
```

#### 2. **Role Folder Explanation**

- **`defaults/`**: Stores default variables for the role.
- **`files/`**: Stores files that need to be copied to the target system.
- **`handlers/`**: Stores handlers that are triggered when certain events occur (e.g., service restart).
- **`meta/`**: Contains metadata for the role, such as dependencies.
- **`tasks/`**: Contains the main tasks for the role (installing Apache2, configuring it, etc.).
- **`templates/`**: Contains Jinja2 templates used to generate configuration files.
- **`vars/`**: Contains variables specific to the role.
- **`tests/`**: Contains tests for the role (optional but recommended).

---

### 3. **Role File Details**

Let’s go through each folder and file and define its contents.

#### **a. `defaults/main.yml`**

The `defaults` folder contains default variables that can be overridden by the user. For example, you might want to make the Apache version configurable.

```yaml
# roles/apache2/defaults/main.yml
---
apache_package: "apache2"
apache_service: "apache2"
apache_port: 80
apache_docroot: "/var/www/html"
```

- **`apache_package`**: The name of the Apache2 package to install.
- **`apache_service`**: The name of the Apache2 service to start and enable.
- **`apache_port`**: The default port for Apache2.
- **`apache_docroot`**: The document root where the website files will be stored.

#### **b. `vars/main.yml`**

The `vars` folder stores role-specific variables. These are variables that are used to customize the behavior of the role. Here, you might define specific configurations for your Apache2 installation.

```yaml
# roles/apache2/vars/main.yml
---
apache_user: "www-data"
apache_group: "www-data"
```

- **`apache_user`**: The user under which Apache runs.
- **`apache_group`**: The group Apache runs under.

#### **c. `tasks/main.yml`**

The `tasks` folder contains the main tasks that define the role’s behavior. In this case, it will contain tasks to install Apache2, start the service, and ensure that the service is enabled.

```yaml
# roles/apache2/tasks/main.yml
---
- name: Install Apache2 package
  apt:
    name: "{{ apache_package }}"
    state: present
    update_cache: yes

- name: Ensure Apache2 service is started and enabled
  service:
    name: "{{ apache_service }}"
    state: started
    enabled: yes

- name: Create the document root directory
  file:
    path: "{{ apache_docroot }}"
    state: directory
    owner: "{{ apache_user }}"
    group: "{{ apache_group }}"
    mode: '0755'

- name: Copy a simple index.html file to the document root
  template:
    src: "index.html.j2"
    dest: "{{ apache_docroot }}/index.html"
    owner: "{{ apache_user }}"
    group: "{{ apache_group }}"
    mode: '0644'
```

- **Task 1**: Installs the Apache2 package using the `apt` module.
- **Task 2**: Ensures the Apache2 service is running and enabled to start on boot.
- **Task 3**: Creates the document root directory `/var/www/html` and ensures it has the right ownership and permissions.
- **Task 4**: Copies an `index.html` file to the Apache2 document root using a Jinja2 template.

#### **d. `handlers/main.yml`**

The `handlers` folder contains handlers, which are tasks that are triggered when notified. For example, you may want to restart Apache after making a configuration change.

```yaml
# roles/apache2/handlers/main.yml
---
- name: Restart Apache service
  service:
    name: "{{ apache_service }}"
    state: restarted
```

- **Handler**: Restarts the Apache service when notified by other tasks.

#### **e. `files/`**

The `files` folder is used to store static files that need to be copied to the target system. For this example, we don’t have any static files to copy, but we could add custom files here if needed (e.g., SSL certificates, logs).

For this example, we'll leave this folder empty.

#### **f. `templates/index.html.j2`**

The `templates` folder is used to store Jinja2 templates. These templates are rendered and copied to the target system. Here, we’ll create a simple `index.html.j2` file for the Apache document root.

```html
<!-- roles/apache2/templates/index.html.j2 -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Apache2</title>
</head>
<body>
    <h1>Apache2 is working!</h1>
    <p>Welcome to your Apache2 server on Ubuntu.</p>
</body>
</html>
```

- This is a simple HTML template rendered with Jinja2, which will be copied to `/var/www/html/index.html` on the target machine.

#### **g. `meta/main.yml`**

The `meta` folder contains metadata about the role, such as role dependencies. In this example, there are no dependencies, so the file will be simple.

```yaml
# roles/apache2/meta/main.yml
---
dependencies: []
```

- **`dependencies`**: This section lists any other roles that should be applied before or after this role. In this case, we don’t have any dependencies.

#### **h. `tests/test.yml`**

The `tests` folder is used to store test playbooks that validate the role. You can write tests to ensure that Apache is installed, the service is running, and the `index.html` file is in place.

```yaml
# roles/apache2/tests/test.yml
---
- name: Test Apache2 installation
  hosts: localhost
  gather_facts: yes

  roles:
    - apache2

  tasks:
    - name: Check if Apache2 is installed
      command: dpkg -l | grep apache2
      register: apache_installed
      failed_when: apache_installed.stdout == ""
      changed_when: false

    - name: Ensure Apache service is running
      service:
        name: "{{ apache_service }}"
        state: started

    - name: Ensure the index.html file exists
      stat:
        path: "{{ apache_docroot }}/index.html"
      register: index_file
      failed_when: index_file.stat.exists == false
```

- **Task 1**: Checks if Apache2 is installed by running `dpkg -l | grep apache2`.
- **Task 2**: Ensures the Apache2 service is running.
- **Task 3**: Ensures that the `index.html` file exists at the document root.

---

### 4. **Running the Role**

To run this role, you need a playbook that applies the role to a target machine. Here’s an example playbook that applies the `apache2` role:

**`site.yml`**:

```yaml
---
- name: Install Apache2 on Ubuntu
  hosts: all
  become: yes
  roles:
    - apache2
```

You can run the playbook with the following command:

```bash
ansible-playbook site.yml -i hosts
```

Make sure to replace `hosts` with your inventory file that defines the target systems.

### 5. **Conclusion**

In this tutorial, we created a comprehensive Ansible role to install and configure Apache2 on an Ubuntu system. We structured the role using the standard Ansible directory layout, with the following key components:

- **defaults**: Default variables for Apache2 configuration.
- **vars**: Role-specific variables, such as the Apache user and group.
- **tasks**: The core tasks, including installing Apache2 and configuring it.
- **handlers**: A handler to restart Apache2 when needed.
- **files**: Static files (empty in this example, but could include SSL certificates, etc.).
- **templates**: Jinja2 templates for generating dynamic configuration files (e.g., `index.html`).
- **meta**: Role metadata, including dependencies.
- **tests**: Test playbooks to ensure the role works correctly.

By using this role, you can easily install and configure Apache2 across multiple Ubuntu servers while maintaining a clean and reusable structure for your Ansible configurations.