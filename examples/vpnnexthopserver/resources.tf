resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
  name        = "tf_vpnnexthopserver"
  nexthopip   = "2.6.1.5"
  nexthopport = "200"
}
