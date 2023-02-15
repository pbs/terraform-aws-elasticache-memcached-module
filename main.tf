resource "aws_elasticache_cluster" "cluster" {
  cluster_id           = local.name
  engine               = local.engine
  node_type            = var.node_type
  num_cache_nodes      = var.nodes
  parameter_group_name = var.parameter_group_name

  apply_immediately = var.apply_immediately

  az_mode                      = var.az_mode
  availability_zone            = var.availability_zone
  preferred_availability_zones = var.preferred_availability_zones

  engine_version         = var.engine_version
  maintenance_window     = var.maintenance_window
  notification_topic_arn = var.notification_topic_arn
  port                   = var.port
  security_group_ids     = local.security_group_ids
  subnet_group_name      = local.subnet_group_name

  tags = local.tags
}
