package instance

import (
	"context"
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_images"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCImageClient
type IBMPPCImageClient struct {
	IBMPPCClient
}

// NewIBMPPCImageClient
func NewIBMPPCImageClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCImageClient {
	return &IBMPPCImageClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get an Image
func (f *IBMPPCImageClient) Get(id string) (*models.Image, error) {
	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetImageOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get Image Operation for image %s", id)
	}
	return resp.Payload, nil
}

// Get All Images that are imported into Power Instance
func (f *IBMPPCImageClient) GetAll() (*models.Images, error) {
	params := p_cloud_images.NewPcloudCloudinstancesImagesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PPC Images of the PVM instance %s : %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PPC Images of the PVM instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Stock Image
func (f *IBMPPCImageClient) Create(body *models.CreateImage) (*models.Image, error) {
	if body.Source == nil || len(*body.Source) == 0 {
		body.Source = core.StringPtr("root-project")
	}
	params := p_cloud_images.NewPcloudCloudinstancesImagesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	respok, respcreated, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateImageOperationFailed, f.cloudInstanceID, err)
	}
	if respok != nil && respok.Payload != nil {
		return respok.Payload, nil
	}
	if respcreated != nil && respcreated.Payload != nil {
		return respcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Image Operation for cloud instance %s", f.cloudInstanceID)
}

// Import an Image
func (f *IBMPPCImageClient) CreateCosImage(body *models.CreateCosImageImportJob) (imageJob *models.JobReference, err error) {
	params := p_cloud_images.NewPcloudV1CloudinstancesCosimagesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	resp, err := f.session.Power.PCloudImages.PcloudV1CloudinstancesCosimagesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform Create COS Image Operation for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create COS Image Operation for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Export an Image
func (f *IBMPPCImageClient) ExportImage(id string, body *models.ExportImage) (*models.JobReference, error) {
	params := p_cloud_images.NewPcloudV2ImagesExportPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithImageID(id).WithBody(body)
	resp, err := f.session.Power.PCloudImages.PcloudV2ImagesExportPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Export COS Image for image id %s with error %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Export COS Image for image id %s", id)
	}
	return resp.Payload, nil
}

// Delete an Image
func (f *IBMPPCImageClient) Delete(id string) error {
	params := p_cloud_images.NewPcloudCloudinstancesImagesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithImageID(id)
	_, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PPC Image %s: %w", id, err)
	}
	return nil
}

// Get a Stock Image
func (f *IBMPPCImageClient) GetStockImage(id string) (*models.Image, error) {
	params := p_cloud_images.NewPcloudCloudinstancesStockimagesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesStockimagesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get PPC Stock Image %s for cloud instance %s: %w", id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PPC Stock Image %s for cloud instance %s", id, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Stock Images
func (f *IBMPPCImageClient) GetAllStockImages(includeSAP bool, includeVTl bool) (*models.Images, error) {
	params := p_cloud_images.NewPcloudCloudinstancesStockimagesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).
		WithSap(&includeSAP).
		WithVtl(&includeVTl)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesStockimagesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get Stock Images with (SAP=%t, VTL=%t) for cloud instance %s: %w", includeSAP, includeVTl, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get Stock Images with (SAP=%t, VTL=%t) for cloud instance %s", includeSAP, includeVTl, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Stock SAP Images
func (f *IBMPPCImageClient) GetAllStockSAPImages() (*models.Images, error) {
	// get stock images. include all available SAP images
	images, err := f.GetAllStockImages(true, false)
	if err != nil {
		return nil, err
	}

	// select SAP images
	sapImages := new(models.Images)
	for _, image := range images.Images {
		if image.Specifications.ImageType == "stock-sap" {
			sapImages.Images = append(sapImages.Images, image)
		}
	}
	return sapImages, nil
}

// Get All Stock VTL Images
func (f *IBMPPCImageClient) GetAllStockVTLImages() (*models.Images, error) {
	// get stock images. include all available stock VTL images
	images, err := f.GetAllStockImages(false, true)
	if err != nil {
		return nil, err
	}

	// select VTL images
	vtlImages := new(models.Images)
	for _, image := range images.Images {
		if image.Specifications.ImageType == "stock-vtl" {
			vtlImages.Images = append(vtlImages.Images, image)
		}
	}
	return vtlImages, nil
}

// Return true if Image is a VTL Image
func (f *IBMPPCImageClient) IsVTLImage(imageId string) (bool, error) {
	images := new(models.Images)

	// get all stock vtl images
	stockVTLImages, err := f.GetAllStockVTLImages()
	if err != nil {
		return false, err
	}
	images.Images = append(images.Images, stockVTLImages.Images...)

	// get all images
	cloudInstanceImages, err := f.GetAll()
	if err != nil {
		return false, err
	}
	images.Images = append(images.Images, cloudInstanceImages.Images...)

	// search for image in list of all images
	for _, image := range images.Images {
		if imageId == *image.ImageID {
			if image.Specifications.ImageType == "stock-vtl" {
				return true, nil
			} else {
				return false, fmt.Errorf("image with id: %s is not a VTL image", imageId)
			}
		}
	}
	return false, fmt.Errorf("image with id: %s is not a VTL image", imageId)
}
