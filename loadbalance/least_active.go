package loadbalance

// LeastActive 最小活跃数负载均衡算法，为每个服务提供者Invoker分配一个active,代表活跃数带下，调用之前做自增操作，
// 调用完成后做自减操作， active 数量越小，就优先分配
func LeastActive() {

}
