  - action: question
    resource: "*"
    effect: deny
  - action: edit
    resource: "*"
    effect: deny
  - action: edit
    resource: ".agent/**"
    effect: allow
  - action: shell
    resource: "*"
    effect: deny
