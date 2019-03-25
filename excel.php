<?php

class GoExcel
{

    const URL = "http://127.0.0.1:8888/v1/sql/Excel";

    private static function https_request($url, $data = NULL)
    {
        $curl = curl_init();
        curl_setopt($curl,CURLOPT_HEADER, false);
        curl_setopt($curl,CURLOPT_POST, true);
        curl_setopt($curl,CURLOPT_BINARYTRANSFER,false);
        $post_data = http_build_query($data);
        curl_setopt($curl, CURLOPT_POSTFIELDS,$post_data);
        curl_setopt($curl, CURLOPT_URL,$url);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);

        $output = curl_exec($curl);
        curl_close($curl);
        return $output;
    }


    /**
     * @param $filename
     * @param $sql
     * @param $dataParm
     * $param = array(
	 *	'ID'=>'id',
	 *	'银行名称'=>'name',
	 *	'电话'=>'phone',
	 *	'地址'=>'addr',
	 *	'银行行号'=>'number',
	 *	'邮编'=>'zip',
     * );
	 * $sql = "select id,name,phone,addr,number,zip from res_bank_list";
	 * GoExcel::get("银行列表" . date("YmdHis") . ".xlsx",$sql,$param);
     */

    public static function get($filename, $sql, $dataParm)
    {

        $data = array();
        $sql = str_replace("`", "", $sql);
        $data['sql'] = $sql;
        //处理特殊的 这样就是排序过的数据
        $i = 0;
        foreach ($dataParm as $key => $value){
                $data[$i] = $key.":".$value;
            $i++;
        }
        $filename = iconv("utf-8", "gb2312", $filename);
        header("Content-Type: application/force-download");
        header("Content-Type: application/octet-stream");
        header("Content-Type: application/download");
        header('Content-Disposition:inline;filename="' . $filename . '"');
        header("Content-Transfer-Encoding: binary");
        header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
        header("Cache-Control: must-revalidate, post-check=0, pre-check=0");
        header("Pragma: no-cache");
        echo self::https_request(self::URL, $data);
    }
}

	$param = array(
		'ID'=>'id',
		'银行名称'=>'name',
		'电话'=>'phone',
		'地址'=>'addr',
		'银行行号'=>'number',
		'邮编'=>'zip',
    );
	$sql = "select id,name,phone,addr,number,zip from res_bank_list";
	GoExcel::get("银行列表" . date("YmdHis") . ".xlsx",$sql,$param);