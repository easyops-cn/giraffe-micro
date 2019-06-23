/*
 * ens_sdk.h
 *
 *  Created on: 2015-12-10
 *      Author: hzp
 */

#ifndef ENS_SDK_H_
#define ENS_SDK_H_

#include <stdint.h>

#define IP_LENGTH 16
#define LICENSE_KEY_LENGTH 32
#define LICENSE_VAL_LENGTH 512

//extern "C"{

/**
 *通过名字拉IP和端口
 *参数： src_name 主调名字
 *		dst_name 被调名字
 *      out_ip 返回的IP
 *      out_ip_len out_ip数组的长度
 *      out_port 返回的端口
 *返回值：session_id 用于上报数据的session_id, 小于0为失败
**/
int64_t get_service_by_name(const char *src_name, const char *dst_name, char out_ip[], const int out_ip_len, int *out_port);

/**
 *通过名字批量拉IP和端口
 *参数： src_name 主调服务名
 *	    dst_name 被调服务名
 *      out_ip_arr 返回的IP数组， 数组大小为arr_size
 *      out_port_arr 返回的端口数组， 数组大小为arr_size
 *      arr_size	入参：传入的数组大小。 出参：实际返回的IP端口个数
 *返回值：session_id 用于上报数据的session_id, 小于0为失败
**/
int64_t get_multi_service_by_name(const char* src_name, const char *dst_name, char out_ip_arr[][IP_LENGTH], int out_port_arr[], int *arr_size);

/**
 * 调用服务是否成功接口1
 * 参数： session_id get_service_by_name返回的session_id
 *		 dst_interface 被调接口
 *       ret_code   接口调用的返回码，0为成功
 *       delay		接口调用延时(毫秒)
 *       code_point	出错代码位置
 *       err_stack	错误堆栈
 * 返回值：小于0为失败
**/
int report_stat(int64_t session_id, const char* dst_interface, int ret_code, int delay, const char* code_point, const char *err_stack);

/**
 * 调用服务是否成功接口2
 * 参数： parent_id 上一级调用的id
 *       request_id   请求id
 *       step_id   步骤id
 *       src_name	主调服务名
 *       dst_name	被调服务名
 *       dst_interface	被调接口名
 *       dst_ip		被调ip
 *       dst_port	被调端口
 *       status		调用状态（0为成功）
 *       delay		调用延时（毫秒）
 *       code_point	代码位置
 *       err_stack	错误堆栈
 * 返回值：小于0为失败
**/
int report_stat_all(const char *parent_id, const char * request_id, int step_id,
		const char *src_name, const char *dst_name, const char *dst_interface,
		const char *dst_ip, int dst_port, int status, int delay, const char* code_point, const char *err_stack);


/**
 * 获取license信息
 *参数： key_list 返回的license key列表
 *	    val_list 返回的license val列表
 *      arr_size	入参：传入的数组大小。 出参：实际返回的key val个数
 *返回值： 小于0为失败
**/
int64_t get_easyops_license_info(char key_list[][LICENSE_KEY_LENGTH], char val_list[][LICENSE_VAL_LENGTH], int *arr_size);

//}

#endif /* ENS_SDK_H_ */
