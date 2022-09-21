output "name" {
  description = "The name of the ElastiCache cluster"
  value       = module.memcached.name
}

output "arn" {
  description = "The ARN of the ElastiCache cluster"
  value       = module.memcached.arn
}

output "sg_ids" {
  description = "The security group ids"
  value       = module.memcached.sg_ids
}


output "cache_nodes" {
  description = "List of node objects including id, address, port and availability_zone"
  value       = module.memcached.cache_nodes
}

output "cluster_address" {
  description = "DNS name of the cache cluster without the port appended"
  value       = module.memcached.cluster_address
}

output "configuration_endpoint" {
  description = "Configuration endpoint to allow host discovery"
  value       = module.memcached.configuration_endpoint
}

output "tags" {
  description = "The tags"
  value       = module.memcached.tags
}

output "lambda_name" {
  description = "The name of the lambda function"
  value       = module.lambda.name
}
