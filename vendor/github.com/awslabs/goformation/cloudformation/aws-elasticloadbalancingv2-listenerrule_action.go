package cloudformation

import (
	"encoding/json"
)

// AWSElasticLoadBalancingV2ListenerRule_Action AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::ListenerRule.Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type AWSElasticLoadBalancingV2ListenerRule_Action struct {

	// TargetGroupArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-targetgrouparn
	TargetGroupArn *Value `json:"TargetGroupArn,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule.Action"
}

func (r *AWSElasticLoadBalancingV2ListenerRule_Action) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
