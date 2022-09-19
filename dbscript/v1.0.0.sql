create table sys_activation_code
(
    id            serial primary key,
    status        boolean default false,
    code          varchar(64),
    activate_time timestamp,
    create_time   timestamp,
    expire_time   timestamp
);

comment on column sys_activation_code.status is '是否激活 默认未激活';

comment on column sys_activation_code.code is '激活码';

comment on column sys_activation_code.activate_time is '激活时间';

comment on column sys_activation_code.create_time is '激活码生成时间';

comment on column sys_activation_code.expire_time is '过期时间';




