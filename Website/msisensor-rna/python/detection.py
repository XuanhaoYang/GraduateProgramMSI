#!/usr/bin/python
# -*- coding: UTF-8 -*-

import os
import sys
import argparse
import joblib
import pandas as pd
from sklearn.preprocessing import StandardScaler
import eli5
import json

curpath = os.path.abspath(os.path.dirname(sys.argv[0]))
sys.path.append(os.path.dirname(curpath))


def args_process():
    """
    argument procress
    """
    parser = argparse.ArgumentParser(
        description="you should add those parameter")  # 这些参数都有默认值，当调用parser.print_help()或者运行程序时由于参数不正确(此时python解释器其实也是调用了pring_help()方法)时，                                                                     # 会打印这些描述信息，一般只需要传递description参数，如上。
    parser.add_argument('--model')
    parser.add_argument('--input')
    parser.add_argument('--output')
    args = parser.parse_args()  # 步骤三
    return args


def main():
    paras = args_process()
    # print("model = " + paras.model)
    # print("input = " + paras.input)
    res = detection(paras.input, paras.model, paras.output)
    print(res)


def zcoreScaler(data):
    x = data
    scaler = StandardScaler()
    x_std = pd.DataFrame(scaler.fit_transform(x.T)).T
    x_std.index = x.index
    x_std.columns = x.columns
    return x_std


def detection(input_path, model_path, output_path):
    model = joblib.load(model_path)
    genes = ['MLH1', 'EPM2AIP1','MSH4','RPL22L1', 'HSPH1','TRIM7', 'PRTFDC1','SMAP1', 'DDX27','CXCL9', 'WDTC1','FECH',  'SFXN1','COL18A1']
    data = pd.read_csv(input_path, index_col=0).fillna(0)
    genes_this = data.columns.to_list()
    if len(set(genes) - set(genes_this)) > 0:
        res = {"code": 0, "msg": ",".join(set(genes) - set(genes_this))}
    else:
        data_scale = zcoreScaler(data[genes])
        y_score = model.predict(data_scale)
        y_score = y_score.round(3)
        cutoff = 6
        y_pre = y_score.copy()
        y_pre[y_pre < cutoff] = 0
        y_pre[y_pre >= cutoff] = 1

        samp = eli5.explain_prediction_df(model, pd.DataFrame(data_scale.values), feature_names=genes)
        samp.index = list(samp["feature"])
        samp = samp.T[genes].T.iloc[:, 2:]
        samp["weight"] = samp["weight"].map(lambda x: abs(x))
        samp["weight"] = samp["weight"] * 100 / sum(samp["weight"])
        samp["value"] = data[genes].values[0]
        samp = samp.round(3)
        res = {"code": 1, 'weight': json.dumps(list(samp.weight)), 'value': json.dumps(list(samp.value)), 'msg': "",
               "score": y_score[0],
               "pre": y_pre[0]}
    res = json.dumps(res)
    # output_file = open(output_path, "w")
    # output_file.write(res)
    # output_file.close()
    return res


if __name__ == "__main__":
    main()
    pass
