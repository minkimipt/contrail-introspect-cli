| purpose | type | description | progress |
|--|--|--|--|
| check_contrail_analytics__status | global/introspec | Check if docker container with contrail-analytics is running | |
| check_contrail_analytics__api | global/introspec | Check contrail analytics API availability | |
| check_contrail_analytics__api_port | global/introspec | Check if port of analytics api is listening | |
| check_contrail_analyticsdb__status | global/introspec | Check if docker container with contrail-analyticsdb is running | |
| check_contrail_controller__status | global/introspec | Check if docker container with contrail-controller is running | |
| check_contrail_controller__api_logs | global/introspec | Check number of error messages in api logs (delta) | |
| check_contrail_controller__api_response_time | global/introspec | Check api response time from external system | |
| check_contrail_controller__schema_logs | global/introspec | Check number of error messages in schema logs (delta) | |
| check_contrail_controller__svc_monitor_logs | global/introspec | Check number of error messages in svc_monitor logs (delta) | |
| check_contrail__vrouter_agent | Global/Introspect | Check contrail-vrouter-agent status | |
| check_contrail__vrouter_nodemgr | Global/Introspect | Check contrail-vrouter-nodemgr status | |
| check_contrail__vrouter_smartnic | Global/Introspect | Check smartnic status using /opt/netronome/libexec/nfp-vrouter-status -r | |
| check_contrail__vrouter_smartnic_metrics | Global/Introspect | Check smartnic metrics using /opt/netronome/libexec/nfp-vr-syscntrs.sh -z | |
| check_contrail__vrouter_xmpp | Global/Introspect | Check number of XMPP sessions | |
| check_contrail__vrouter_cpu | Global/Introspect | Check contrail-vrouter-agent cpu usage | done |
| check_contrail__vrouter_memory | Global/Introspect | Check contrail-vrouter-agent memory usage | done |
| check_contrail__vrouter_flaps | Global/Introspect | Check contrail-vrouter-agent for number of flaps | |
| check_contrail__dropstats_invalid_source | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_source entry | |
| check_contrail__dropstats_flow_queue_limit_exceeded | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_queue_limit_exceeded entry | |
| check_contrail__dropstats_drop_new_flow | Global/Introspect | Check dropstats for check_contrail__dropstats_drop_new_flow entry | |
| check_contrail__dropstats_invalid_nh | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_nh entry | |
| check_contrail__dropstats_interface_tx_discard | Global/Introspect | Check dropstats for check_contrail__dropstats_interface_tx_discard entry | |
| check_contrail__dropstats_flow_table_full | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_table_full entry | |
| check_contrail__dropstats_active_flows | Global/Introspect | Check dropstats for check_contrail__dropstats_active_flows entry | |
| check_contrail__dropstats_new_flows_setup_rate | Global/Introspect | Check dropstats for check_contrail__dropstats_new_flows_setup_rate entry | |
| check_contrail__dropstats_overlay_connectivity | Global/Introspect | Check dropstats for check_contrail__dropstats_overlay_connectivity entry | |
| check_contrail__dropstats_flow_drop_due_to_max_limit | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_drop_due_to_max_limit entry | |
| check_contrail__dropstats_flow_drop_due_to_linklocal_limit | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_drop_due_to_linklocal_limit entry | |
| check_contrail__dropstats_flow_drop_due_to_linklocal_limit | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_drop_due_to_linklocal_limit entry | |
| check_contrail__dropstats_flow_export_disable_drops | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_export_disable_drops entry | |
| check_contrail__dropstats_flow_export_drops | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_export_drops entry | |
| check_contrail__dropstats_rid | Global/Introspect | Check dropstats for check_contrail__dropstats_rid entry | |
| check_contrail__dropstats_discard | Global/Introspect | Check dropstats for check_contrail__dropstats_discard entry | |
| check_contrail__dropstats_invalid_if | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_if entry | |
| check_contrail__dropstats_arp_not_me | Global/Introspect | Check dropstats for check_contrail__dropstats_arp_not_me entry | |
| check_contrail__dropstats_invalid_arp | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_arp entry | |
| check_contrail__dropstats_trap_no_if | Global/Introspect | Check dropstats for check_contrail__dropstats_trap_no_if entry | |
| check_contrail__dropstats_nowhere_to_go | Global/Introspect | Check dropstats for check_contrail__dropstats_nowhere_to_go entry | |
| check_contrail__dropstats_flow_no_memory | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_no_memory entry | |
| check_contrail__dropstats_flow_invalid_protocol | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_invalid_protocol entry | |
| check_contrail__dropstats_flow_nat_no_rflow | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_nat_no_rflow entry | |
| check_contrail__dropstats_flow_action_drop | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_action_drop entry | |
| check_contrail__dropstats_flow_action_invalid | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_action_invalid entry | |
| check_contrail__dropstats_flow_unusable | Global/Introspect | Check dropstats for check_contrail__dropstats_flow_unusable entry | |
| check_contrail__dropstats_interface_drop | Global/Introspect | Check dropstats for check_contrail__dropstats_interface_drop entry | |
| check_contrail__dropstats_ttl_exceeded | Global/Introspect | Check dropstats for check_contrail__dropstats_ttl_exceeded entry | |
| check_contrail__dropstats_invalid_label | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_label entry | |
| check_contrail__dropstats_invalid_protocol | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_protocol entry | |
| check_contrail__dropstats_interface_rx_discard | Global/Introspect | Check dropstats for check_contrail__dropstats_interface_rx_discard entry | |
| check_contrail__dropstats_invalid_mcast_source | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_mcast_source entry | |
| check_contrail__dropstats_head_alloc_fail | Global/Introspect | Check dropstats for check_contrail__dropstats_head_alloc_fail entry | |
| check_contrail__dropstats_mcast_clone_fail | Global/Introspect | Check dropstats for check_contrail__dropstats_mcast_clone_fail entry | |
| check_contrail__dropstats_rewrite_fail | Global/Introspect | Check dropstats for check_contrail__dropstats_rewrite_fail entry | |
| check_contrail__dropstats_misc | Global/Introspect | Check dropstats for check_contrail__dropstats_misc entry | |
| check_contrail__dropstats_invalid_packet | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_packet entry | |
| check_contrail__dropstats_cksum_err | Global/Introspect | Check dropstats for check_contrail__dropstats_cksum_err entry | |
| check_contrail__dropstats_no_fmd | Global/Introspect | Check dropstats for check_contrail__dropstats_no_fmd entry | |
| check_contrail__dropstats_invalid_vnid | Global/Introspect | Check dropstats for check_contrail__dropstats_invalid_vnid entry | |
| check_contrail__dropstats_frag_err | Global/Introspect | Check dropstats for check_contrail__dropstats_frag_err entry | |
| check_contrail__dropstats_no_memory | Global/Introspect | Check dropstats for check_contrail__dropstats_no_memory entry | |
