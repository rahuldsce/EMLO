# Initiating the training process
trainer = Trainer(
    module_file="cifar10_train.py",
    data_module_file="cifar10_datamodule.py",
    module_file_args=args,
    data_module_args=data_module_args,
    trainer_args=trainer_args
    #unit test if True runs once or specifify int, use debug training/validation loop, use test if everything is working fine
    fast_dev_run=True
    precision=16,
    accelerator="dpp",
    #overfit on small portion of data, if doesn't overfit then it won't work with large datasets
    overfit_batches=0.01
)