Using **environment variables** in **Ansible** playbooks is essential when you want to make your automation more flexible, secure, and environment-aware. Environment variables are typically used to provide dynamic values like API keys, paths, or other configuration settings that might vary depending on the environment (e.g., development, staging, production) or runtime.

In Ansible, environment variables can be accessed and managed in several ways, including through **tasks**, **playbooks**, and **roles**. You can also pass environment variables to a command or script that Ansible runs.

### Key Concepts for Using Environment Variables in Ansible

1. **Accessing Environment Variables**: You can access environment variables in Ansible playbooks using the `lookup` plugin or `env` lookup.

2. **Setting Environment Variables for Tasks**: You can explicitly set environment variables for individual tasks to ensure a command or script runs with specific environment settings.

3. **Passing Environment Variables to Commands**: Ansible allows you to pass environment variables to commands, modules, and scripts during execution.

4. **Using Environment Variables in Playbooks or Roles**: You can define environment variables that should be used during the playbook execution or by tasks.

---

### Accessing Environment Variables in Ansible

#### 1. **Using `env` Lookup Plugin**

The simplest way to access environment variables is by using the **`env` lookup** plugin. The `env` plugin allows you to read the current environment variables during playbook execution.

**Example**: Accessing an environment variable like `HOME`

```yaml
- name: Example playbook to access environment variables
  hosts: localhost
  tasks:
    - name: Get HOME environment variable
      debug:
        msg: "The HOME directory is {{ lookup('env', 'HOME') }}"
```

**Explanation**:
- The `lookup('env', 'HOME')` reads the value of the `HOME` environment variable and makes it available in the playbook.
- This example will print the path of the `HOME` directory for the user running the playbook.

**Example Output**:
```yaml
TASK [Get HOME environment variable] ***
ok: [localhost] => (item=None) =>
  msg: "The HOME directory is /home/ansible"
```

#### 2. **Using `env` in `vars`**

You can also use environment variables in the `vars` section to define variables dynamically based on the current environment.

**Example**: Set a variable based on an environment variable

```yaml
- name: Example playbook with env variables in vars
  hosts: localhost
  vars:
    user_home: "{{ lookup('env', 'HOME') }}"
  tasks:
    - name: Show the user's home directory
      debug:
        msg: "The user home directory is {{ user_home }}"
```

**Explanation**:
- The variable `user_home` is dynamically set to the value of the `HOME` environment variable, which is then used in the playbook.

---

### Setting Environment Variables for Specific Tasks

Sometimes you need to set environment variables for specific tasks rather than globally for the entire playbook. Ansible allows you to set environment variables specifically for tasks or commands using the `environment` directive.

#### Example 1: Setting Environment Variables for a Task

```yaml
- name: Example playbook to set environment variables for a task
  hosts: localhost
  tasks:
    - name: Run a command with a specific environment variable
      command: echo $MY_VAR
      environment:
        MY_VAR: "Hello, World!"
```

**Explanation**:
- The task runs the `echo $MY_VAR` command, but with `MY_VAR` set to "Hello, World!" for this task only.
- This environment variable is only available to the task it is defined in, not the entire playbook.

**Example Output**:
```yaml
TASK [Run a command with a specific environment variable] ***
changed: [localhost] => (item=None) =>
  msg: "Hello, World!"
```

#### Example 2: Passing Environment Variables to Shell Scripts

```yaml
- name: Run a shell script with environment variables
  hosts: localhost
  tasks:
    - name: Execute shell script with environment variables
      shell: |
        echo "Script started"
        echo "The API key is $API_KEY"
      environment:
        API_KEY: "12345"
```

**Explanation**:
- In this case, the `API_KEY` environment variable is passed to the shell script. The `echo` commands will output the value of the `API_KEY` during script execution.

**Example Output**:
```yaml
TASK [Execute shell script with environment variables] ***
changed: [localhost] => (item=None) =>
  msg: |
    Script started
    The API key is 12345
```

---

### Using Environment Variables with Ansible Roles

If you're working with **roles** and want to set environment variables for tasks in a role, you can use the `environment` directive in the role’s task files or define them in the role’s `defaults/main.yml` or `vars/main.yml` if they are global for that role.

**Example**: Set environment variables in a role task

```yaml
# roles/myrole/tasks/main.yml
- name: Run a task with environment variables
  shell: |
    echo "Using MY_ENV_VAR: $MY_ENV_VAR"
  environment:
    MY_ENV_VAR: "{{ my_role_variable }}"
```

**Explanation**:
- The environment variable `MY_ENV_VAR` is passed from the role's `my_role_variable` value.

---

### Using Environment Variables in Templates

You can also use environment variables inside Ansible **templates** (Jinja2 templates) to generate configuration files or perform dynamic substitutions.

#### Example: Using Environment Variables in a Template

```yaml
- name: Example playbook using environment variable in template
  hosts: localhost
  tasks:
    - name: Create a config file from template
      template:
        src: "/path/to/my_config.j2"
        dest: "/tmp/my_config.conf"
      environment:
        MY_VAR: "Hello from environment"
```

**Template (`my_config.j2`)**:
```jinja
# Configuration file
env_var_value={{ lookup('env', 'MY_VAR') }}
```

**Explanation**:
- In this example, the environment variable `MY_VAR` is accessed inside the template and written to a file.
- The template file uses the Jinja2 `lookup` function to access `MY_VAR`.

---

### Using Environment Variables with `ansible-playbook`

You can also pass environment variables to an Ansible playbook at runtime using the `-e` (extra-vars) option or by setting them directly in the shell.

#### Example 1: Passing Environment Variables with `ansible-playbook`

```bash
MY_VAR=some_value ansible-playbook -e "some_other_var=another_value" playbook.yml
```

**Explanation**:
- The environment variable `MY_VAR` is passed to the playbook as an environment variable. You can use this variable inside the playbook via the `lookup('env', 'MY_VAR')` function.

#### Example 2: Accessing Environment Variables in `ansible-playbook`

```bash
ansible-playbook playbook.yml
```

Inside the playbook:

```yaml
- name: Use environment variable
  debug:
    msg: "The environment variable MY_VAR is {{ lookup('env', 'MY_VAR') }}"
```

**Explanation**:
- The `MY_VAR` environment variable is accessed inside the playbook using `lookup('env', 'MY_VAR')`.

---

### Best Practices for Using Environment Variables in Ansible

1. **Security**:
   - Be careful when using environment variables for sensitive information like API keys, passwords, or tokens. Consider using **Ansible Vault** to encrypt sensitive values.
   - Avoid hardcoding sensitive values directly in playbooks or inventory files.

2. **Flexibility**:
   - Use environment variables to configure values that may vary between different environments (e.g., production, staging, development).
   - Using environment variables makes your playbooks more flexible and environment-agnostic.

3. **Documentation**:
   - Always document the environment variables used in your playbooks or roles, especially if they have dependencies or specific values that must be set in the environment.
   
4. **Default Values**:
   - Use the `default` filter when accessing environment variables to provide default values in case they are not set in the environment.

   ```yaml
   my_var: "{{ lookup('env', 'MY_VAR') | default('default_value') }}"
   ```

---

### Conclusion

Using **environment variables** in Ansible is an excellent way to make your playbooks more dynamic, environment-specific, and secure. Whether you're accessing existing system variables, passing custom variables to tasks or scripts, or defining environment-specific configuration values in roles, Ansible provides several powerful ways to handle environment variables.

By leveraging environment variables, you can:
- Easily customize playbook behavior without changing code.
- Improve security by storing sensitive data outside playbooks.
- Build more flexible, portable automation workflows.

The use of environment variables helps you write clean, modular, and maintainable automation that can adapt to different environments with minimal changes.