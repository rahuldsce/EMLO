# Assignment 8 -  Pytorch Lightning Feature

```
trainer = Trainer(
    module_file="cifar10_train.py",
    data_module_file="cifar10_datamodule.py",
    module_file_args=args,
    data_module_args=data_module_args,
    trainer_args=trainer_args
    fast_dev_run=True
    precision=16,
    accelerator="dp",
    overfit_batches=0.01
)
```

## Unit Test
```
fast_dev_run = True/int; 
```
use to debug training/validation loop to check if the model is working fine and free of bugs

## Mixed Precision
```
precision=16
```
Helps to reduce memory footprint during model training resulting in improved performance and faster training

## Accelerator
```
accelerator='dpp'
```
Distributing data over multiple gpu, to speed up training, dp stands for data parallel where the gpu are on the same machine using single disk, whereas dpp stands for distributed data parallel multiple gpu's across multiple machine each with there own disk

### Overfit
```
overfit_batches=0.01
```
use to overfit small amount of data to check if model can classify classes, if it doesn't work on small data then it indicates model will not work on large datasets
