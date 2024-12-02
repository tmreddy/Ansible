In Ansible, "when" conditions allow you to control when a task, block, or role is executed based on the evaluation of an expression. It's a powerful feature to ensure that certain tasks are only run when specific conditions are met, allowing for more flexible, dynamic, and efficient playbooks.

### 1. **Basic Syntax of `when`**
The `when` condition is written as a Jinja2 expression. If the expression evaluates to `True`, the task will run; if `False`, the task will be skipped.

```yaml
- name: Install nginx if not already installed
  ansible.builtin.yum:
    name: nginx
    state: present
  when: ansible_facts.packages.nginx is not defined
```

In the above example, the task will install Nginx only if it is not already present (i.e., not installed).

### 2. **Using Variables in `when`**
You can use variables defined earlier in the playbook or gathered facts in the `when` condition. 

#### Example 1: Checking a boolean variable
```yaml
- name: Install Apache if the install_apache variable is true
  ansible.builtin.yum:
    name: httpd
    state: present
  when: install_apache | default(false)
```

In this example, the task will only execute if the variable `install_apache` is `True`. If the variable isn't defined, it defaults to `false`, and the task will be skipped.

#### Example 2: Comparing values
```yaml
- name: Check if disk space is available
  ansible.builtin.shell: df -h / | tail -n 1
  register: disk_check
  when: disk_check.stdout.find("100%") == -1
```

This example will run a shell command to check disk space, but only if the disk is not 100% full. If `disk_check.stdout` contains "100%", the task will be skipped.

### 3. **Combining Multiple Conditions**
You can combine multiple conditions using `and`, `or`, or parentheses for grouping to create more complex logic.

#### Example 1: Using `and`
```yaml
- name: Install nginx if the system is a debian-based distribution and the user wants nginx
  ansible.builtin.yum:
    name: nginx
    state: present
  when: ansible_facts.os_family == "Debian" and install_nginx == true
```
In this example, the task will only run if both conditions are met: the OS is a Debian family (like Ubuntu) and the `install_nginx` variable is `True`.

#### Example 2: Using `or`
```yaml
- name: Install nginx if either Apache or Nginx is not installed
  ansible.builtin.yum:
    name: nginx
    state: present
  when: ansible_facts.packages.apache2 is not defined or ansible_facts.packages.nginx is not defined
```

This example checks if either Apache or Nginx is missing and installs Nginx if either condition is true.

#### Example 3: Parentheses for grouping
```yaml
- name: Install nginx or apache if the user wants web server
  ansible.builtin.yum:
    name: "{{ item }}"
    state: present
  loop:
    - nginx
    - httpd
  when: (web_server == 'nginx' or web_server == 'apache') and ansible_facts.os_family == 'RedHat'
```

Here, the task will install either Nginx or Apache based on the value of the `web_server` variable, but only if the system is based on Red Hat (e.g., CentOS or RHEL).

### 4. **Using `when` with Loops**
You can also apply `when` conditions within loops. This allows you to perform specific actions only for certain items in a list.

#### Example 1: Applying `when` to each loop item
```yaml
- name: Install multiple packages only if they're not already installed
  ansible.builtin.yum:
    name: "{{ item }}"
    state: present
  loop:
    - nginx
    - httpd
    - mysql
  when: item != "httpd"  # Skip installing httpd
```

This will install all packages in the loop except for `httpd` (Apache).

#### Example 2: Using `when` with `ansible_facts`
```yaml
- name: Install specific packages based on OS family
  ansible.builtin.yum:
    name: "{{ item }}"
    state: present
  loop:
    - nginx
    - apache2
  when: ansible_facts.os_family == 'Debian' and item == 'apache2'
```

Here, `apache2` will only be installed if the OS is Debian-based (e.g., Ubuntu), and the loop will skip the installation of Nginx.

### 5. **Using `when` with Blocks**
`when` can also be applied to entire blocks of tasks. This allows for better organization of related tasks, ensuring that the block is only executed if the condition is true.

#### Example 1: Applying `when` to a block
```yaml
- name: Install web server packages
  block:
    - name: Install nginx
      ansible.builtin.yum:
        name: nginx
        state: present
    - name: Install httpd
      ansible.builtin.yum:
        name: httpd
        state: present
  when: install_web_server | default(false)
```

In this case, both Nginx and Apache will only be installed if the variable `install_web_server` is `True`.

### 6. **Negating a `when` Condition**
You can also negate a condition using `not`.

#### Example 1: Negating a condition
```yaml
- name: Skip task if the user is root
  ansible.builtin.debug:
    msg: "This task will be skipped if the user is root."
  when: ansible_facts.user != "root"
```

This task will only run if the current user is not `root`. The task is skipped when the user is `root`.

### 7. **`when` with Facts and Gathered Information**
You can use facts that are gathered at the start of the playbook to make conditional decisions about whether to run a task or not. Facts provide a wealth of information about the system, such as the operating system, IP addresses, memory, CPU, etc.

#### Example 1: Use facts to conditionally apply tasks
```yaml
- name: Check if the system is Ubuntu
  ansible.builtin.debug:
    msg: "This system is Ubuntu!"
  when: ansible_facts.distribution == "Ubuntu"
```

### 8. **`when` with `register`**
You can use `register` to store the result of a task and then base subsequent decisions on that result.

#### Example 1: Using `register` with `when`
```yaml
- name: Check disk space
  ansible.builtin.shell: df -h / | grep "/dev"
  register: disk_check

- name: Alert if disk is almost full
  ansible.builtin.debug:
    msg: "Disk space is critically low!"
  when: disk_check.stdout.find("100%") != -1
```

In this case, the second task will be executed only if the disk is 100% full, based on the value stored in `disk_check`.

### 9. **Common Pitfalls**
- **Syntax Errors**: Always remember to use correct Jinja2 syntax. For example, using `is defined` to check if a variable exists.
- **Falsy Values**: Be aware that `None`, `False`, `0`, and empty strings `""` are considered falsy in Ansible's Jinja2 expressions.
  
  ```yaml
  when: variable is not none  # Valid check for non-None values
  ```

### Conclusion
The `when` condition in Ansible is a critical feature for building intelligent and flexible playbooks. By using `when`, you can create more dynamic tasks that run only under the right conditions, based on variables, facts, or the results of previous tasks. Combining these with loops, blocks, and complex conditional logic allows you to build sophisticated automation workflows that are both efficient and adaptable.