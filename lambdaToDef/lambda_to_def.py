def convert_lambda_to_def(string):
    split_def = string.split()

    def_string = f'def {split_def[0]}({split_def[3][:-1]}):\n    return {" ".join(split_def[4:])}'

    return def_string
