package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAWSWafWebAcl_basic(t *testing.T) {
	var v waf.WebACL
	wafAclName := fmt.Sprintf("wafacl%s", acctest.RandString(5))
	resourceName := "aws_waf_web_acl.waf_acl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafWebAclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafWebAclConfig(wafAclName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &v),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.4234791575.type", "ALLOW"),
					resource.TestCheckResourceAttr(
						resourceName, "name", wafAclName),
					resource.TestCheckResourceAttr(
						resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "metric_name", wafAclName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSWafWebAcl_group(t *testing.T) {
	var v waf.WebACL
	wafAclName := fmt.Sprintf("wafaclgroup%s", acctest.RandString(5))
	resourceName := "aws_waf_web_acl.waf_acl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafWebAclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSWafWebAclGroupConfig(wafAclName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &v),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.4234791575.type", "ALLOW"),
					resource.TestCheckResourceAttr(
						resourceName, "name", wafAclName),
					resource.TestCheckResourceAttr(
						resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "metric_name", wafAclName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSWafWebAcl_changeNameForceNew(t *testing.T) {
	var before, after waf.WebACL
	wafAclName := fmt.Sprintf("wafacl%s", acctest.RandString(5))
	wafAclNewName := fmt.Sprintf("wafacl%s", acctest.RandString(5))
	resourceName := "aws_waf_web_acl.waf_acl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafWebAclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafWebAclConfig(wafAclName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &before),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.4234791575.type", "ALLOW"),
					resource.TestCheckResourceAttr(
						resourceName, "name", wafAclName),
					resource.TestCheckResourceAttr(
						resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "metric_name", wafAclName),
				),
			},
			{
				Config: testAccAWSWafWebAclConfigChangeName(wafAclNewName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &after),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.4234791575.type", "ALLOW"),
					resource.TestCheckResourceAttr(
						resourceName, "name", wafAclNewName),
					resource.TestCheckResourceAttr(
						resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "metric_name", wafAclNewName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSWafWebAcl_changeDefaultAction(t *testing.T) {
	var before, after waf.WebACL
	wafAclName := fmt.Sprintf("wafacl%s", acctest.RandString(5))
	wafAclNewName := fmt.Sprintf("wafacl%s", acctest.RandString(5))
	resourceName := "aws_waf_web_acl.waf_acl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafWebAclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafWebAclConfig(wafAclName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &before),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.4234791575.type", "ALLOW"),
					resource.TestCheckResourceAttr(
						resourceName, "name", wafAclName),
					resource.TestCheckResourceAttr(
						resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "metric_name", wafAclName),
				),
			},
			{
				Config: testAccAWSWafWebAclConfigDefaultAction(wafAclNewName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &after),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "default_action.2267395054.type", "BLOCK"),
					resource.TestCheckResourceAttr(
						resourceName, "name", wafAclNewName),
					resource.TestCheckResourceAttr(
						resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(
						resourceName, "metric_name", wafAclNewName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSWafWebAcl_disappears(t *testing.T) {
	var v waf.WebACL
	wafAclName := fmt.Sprintf("wafacl%s", acctest.RandString(5))
	resourceName := "aws_waf_web_acl.waf_acl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafWebAclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafWebAclConfig(wafAclName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafWebAclExists(resourceName, &v),
					testAccCheckAWSWafWebAclDisappears(&v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckAWSWafWebAclDisappears(v *waf.WebACL) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*AWSClient).wafconn

		wr := newWafRetryer(conn, "global")
		_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
			req := &waf.UpdateWebACLInput{
				ChangeToken: token,
				WebACLId:    v.WebACLId,
			}

			for _, ActivatedRule := range v.Rules {
				WebACLUpdate := &waf.WebACLUpdate{
					Action: aws.String("DELETE"),
					ActivatedRule: &waf.ActivatedRule{
						Priority: ActivatedRule.Priority,
						RuleId:   ActivatedRule.RuleId,
						Action:   ActivatedRule.Action,
					},
				}
				req.Updates = append(req.Updates, WebACLUpdate)
			}

			return conn.UpdateWebACL(req)
		})
		if err != nil {
			return fmt.Errorf("Error Updating WAF ACL: %s", err)
		}

		_, err = wr.RetryWithToken(func(token *string) (interface{}, error) {
			opts := &waf.DeleteWebACLInput{
				ChangeToken: token,
				WebACLId:    v.WebACLId,
			}
			return conn.DeleteWebACL(opts)
		})
		if err != nil {
			return fmt.Errorf("Error Deleting WAF ACL: %s", err)
		}
		return nil
	}
}

func testAccCheckAWSWafWebAclDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_waf_web_acl" {
			continue
		}

		conn := testAccProvider.Meta().(*AWSClient).wafconn
		resp, err := conn.GetWebACL(
			&waf.GetWebACLInput{
				WebACLId: aws.String(rs.Primary.ID),
			})

		if err == nil {
			if *resp.WebACL.WebACLId == rs.Primary.ID {
				return fmt.Errorf("WebACL %s still exists", rs.Primary.ID)
			}
		}

		// Return nil if the WebACL is already destroyed
		if isAWSErr(err, waf.ErrCodeNonexistentItemException, "") {
			continue
		}

		return err
	}

	return nil
}

func testAccCheckAWSWafWebAclExists(n string, v *waf.WebACL) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebACL ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).wafconn
		resp, err := conn.GetWebACL(&waf.GetWebACLInput{
			WebACLId: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if *resp.WebACL.WebACLId == rs.Primary.ID {
			*v = *resp.WebACL
			return nil
		}

		return fmt.Errorf("WebACL (%s) not found", rs.Primary.ID)
	}
}

func testAccAWSWafWebAclConfig(name string) string {
	return fmt.Sprintf(`resource "aws_waf_ipset" "ipset" {
  name = "%s"
  ip_set_descriptors {
    type = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_waf_rule" "wafrule" {
  depends_on = ["aws_waf_ipset.ipset"]
  name = "%s"
  metric_name = "%s"
  predicates {
    data_id = "${aws_waf_ipset.ipset.id}"
    negated = false
    type = "IPMatch"
  }
}
resource "aws_waf_web_acl" "waf_acl" {
  depends_on = ["aws_waf_ipset.ipset", "aws_waf_rule.wafrule"]
  name = "%s"
  metric_name = "%s"
  default_action {
    type = "ALLOW"
  }
  rules {
    action {
       type = "BLOCK"
    }
    priority = 1 
    rule_id = "${aws_waf_rule.wafrule.id}"
  }
}`, name, name, name, name, name)
}

func testAccAWSWafWebAclGroupConfig(name string) string {
	return fmt.Sprintf(`

resource "aws_waf_rule_group" "wafrulegroup" {
  name = "%s"
  metric_name = "%s"
}
resource "aws_waf_web_acl" "waf_acl" {
  depends_on = ["aws_waf_rule_group.wafrulegroup"]
  name = "%s"
  metric_name = "%s"
  default_action {
    type = "ALLOW"
  }
  rules {
    override_action {
       type = "NONE"
    }
    type = "GROUP"
    priority = 1
    rule_id = "${aws_waf_rule_group.wafrulegroup.id}"
  }
}`, name, name, name, name)
}

func testAccAWSWafWebAclConfigChangeName(name string) string {
	return fmt.Sprintf(`resource "aws_waf_ipset" "ipset" {
  name = "%s"
  ip_set_descriptors {
    type = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_waf_rule" "wafrule" {
  depends_on = ["aws_waf_ipset.ipset"]
  name = "%s"
  metric_name = "%s"
  predicates {
    data_id = "${aws_waf_ipset.ipset.id}"
    negated = false
    type = "IPMatch"
  }
}
resource "aws_waf_web_acl" "waf_acl" {
  depends_on = ["aws_waf_ipset.ipset", "aws_waf_rule.wafrule"]
  name = "%s"
  metric_name = "%s"
  default_action {
    type = "ALLOW"
  }
  rules {
    action {
       type = "BLOCK"
    }
    priority = 1 
    rule_id = "${aws_waf_rule.wafrule.id}"
  }
}`, name, name, name, name, name)
}

func testAccAWSWafWebAclConfigDefaultAction(name string) string {
	return fmt.Sprintf(`resource "aws_waf_ipset" "ipset" {
  name = "%s"
  ip_set_descriptors {
    type = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_waf_rule" "wafrule" {
  depends_on = ["aws_waf_ipset.ipset"]
  name = "%s"
  metric_name = "%s"
  predicates {
    data_id = "${aws_waf_ipset.ipset.id}"
    negated = false
    type = "IPMatch"
  }
}
resource "aws_waf_web_acl" "waf_acl" {
  depends_on = ["aws_waf_ipset.ipset", "aws_waf_rule.wafrule"]
  name = "%s"
  metric_name = "%s"
  default_action {
    type = "BLOCK"
  }
  rules {
    action {
       type = "BLOCK"
    }
    priority = 1 
    rule_id = "${aws_waf_rule.wafrule.id}"
  }
}`, name, name, name, name, name)
}
