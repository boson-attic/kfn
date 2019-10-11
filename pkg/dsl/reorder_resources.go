package dsl

//func ReorderResources(resources []interface{}) []interface{} {
//	sort.Slice(resources, func(i, j int) bool {
//		iR := resources[i]
//		jR := resources[j]
//		switch iR.(type) {
//		case serving_v1.Service:
//			return true
//		case sources_v1_alpha1.CronJobSource:
//			return true
//		}
//		switch jR.(type) {
//		case messaging_v1_alpha1.Subscription:
//			return false
//		}
//		return false
//	})
//	return resources
//}
