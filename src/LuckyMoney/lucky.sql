CREATE TABLE `lucky_account` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '账户编号',
  `account_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '账户名称',
  `account_type` tinyint(2) DEFAULT NULL COMMENT '账户类型',
  `currency_code` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '货币类型编码',
  `user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户编号',
  `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名称',
  `balance` decimal(30,6) DEFAULT NULL COMMENT '账户可用余额',
  `status` tinyint(2) DEFAULT NULL COMMENT '账户状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `lucky_account_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `trade_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易单号',
  `log_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '流水编号',
  `account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '账户编号',
  `target_account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '目标账户编号',
  `user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户编号',
  `target_user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '目标用户编号',
  `amount` decimal(30,6) DEFAULT NULL COMMENT '交易金额',
  `balance` decimal(30,6) DEFAULT NULL COMMENT '交易后余额',
  `change_type` tinyint(2) DEFAULT NULL COMMENT '流水交易类型',
  `change_flag` tinyint(2) DEFAULT NULL COMMENT '交易变化标识',
  `status` tinyint(2) DEFAULT NULL COMMENT '交易状态',
  `decs` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易描述',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `lucky_red_envelope_goods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `envelope_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红包编号',
  `envelope_type` tinyint(2) DEFAULT NULL COMMENT '红包类型',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名称',
  `user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户编号',
  `blessing` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '祝福语',
  `amount` decimal(30,6) DEFAULT NULL COMMENT '红包总金额',
  `amount_one` decimal(30,6) DEFAULT NULL COMMENT '单个红包金额',
  `quantity` int(10) DEFAULT NULL COMMENT '红包总数量',
  `remain_amount` decimal(30,6) DEFAULT NULL COMMENT '红包剩余金额',
  `remain_quantity` int(10) DEFAULT NULL COMMENT '红包剩余数量',
  `expired_at` datetime(3) DEFAULT NULL COMMENT '过期时间',
  `status` tinyint(2) DEFAULT NULL COMMENT '红包状态',
  `order_type` tinyint(2) DEFAULT NULL COMMENT '订单类型',
  `pay_status` tinyint(2) DEFAULT NULL COMMENT '支付状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `lucky_red_envelope_item` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `item_no` bigint(20) DEFAULT NULL COMMENT '详情编号',
  `envelope_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '红包编号',
  `recv_username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名称',
  `recv_user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户编号',
  `amount` decimal(30,6) DEFAULT NULL COMMENT '收到金额',
  `quantity` int(10) DEFAULT NULL COMMENT '收到数量',
  `remain_amount` decimal(30,6) DEFAULT NULL COMMENT '红包剩余金额',
  `account_no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '账户ID',
  `pay_status` tinyint(2) DEFAULT NULL COMMENT '支付状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;