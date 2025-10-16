terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  # Region auto-detected from AWS CLI config or AWS_REGION env var
}

data "aws_region" "current" {}

module "ecr" {
  source = "../modules/ecr"

  project_name = var.project_name
  environment  = var.environment
}

module "networking" {
  source = "../modules/networking"
}

module "security" {
  source = "../modules/security"

  project_name = var.project_name
  environment  = var.environment
  vpc_id       = module.networking.vpc_id
}

module "iam" {
  source = "../modules/iam"
}

module "alb" {
  source = "../modules/alb"

  project_name          = var.project_name
  environment           = var.environment
  vpc_id                = module.networking.vpc_id
  subnet_ids            = module.networking.public_subnet_ids
  alb_security_group_id = module.security.alb_security_group_id
}

module "ecs" {
  source = "../modules/ecs"

  project_name       = var.project_name
  environment        = var.environment
  aws_region         = data.aws_region.current.name
  ecr_repository_url = module.ecr.repository_url
  subnet_ids         = module.networking.public_subnet_ids
  security_group_id  = module.security.ecs_security_group_id
  execution_role_arn = module.iam.ecs_task_execution_role_arn
  task_role_arn      = module.iam.ecs_task_role_arn
  task_cpu           = var.task_cpu
  task_memory        = var.task_memory
  target_group_arn   = module.alb.target_group_arn

  # Auto Scaling configuration
  min_capacity       = var.min_capacity
  max_capacity       = var.max_capacity
  cpu_target_value   = var.cpu_target_value
  scale_in_cooldown  = var.scale_in_cooldown
  scale_out_cooldown = var.scale_out_cooldown
}
