USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_GetServiceOFferingsByIDOrApikey]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_GetServiceOFferingsByIDOrApikey')
BEGIN
    DROP PROCEDURE [Customer].[usp_GetServiceOFferingsByIDOrApikey]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_GetServiceOFferingsByIDOrApikey]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-05-14  1.0           Naveen Prabhaker

TEST CODE 
-- EXEC [Customer].[usp_GetServiceOFferingsByIDOrApikey] @apiKey = 'ZJBdn&6ckT&!2S(K4Qc^vJSyn'

Initial Version

*****************************************************************************/
 

CREATE PROCEDURE [Customer].[usp_GetServiceOFferingsByIDOrApikey] (
@apiKey VARCHAR(500) = '',
@CustomerID INT = 0
)
AS
BEGIN

	SELECT 
		Distinct
		[ServiceOfferingID] = SO.ServiceOfferingID, 
		[Name] = SO.Name, 
		[Description] = ISNULL(SO.ServiceDescription, SO.Name), 
		[InternalReferenceID] = ISNULL(SO.InternalServiceRecordID,0)
	FROM 
		Customer.CustomerCredentials CC
	INNER JOIN Customer C
		ON C.CustomerID = CC.CustomerID
	INNER JOIN Customer.CustomerServiceOfferings CSO
		ON CSO.CustomerID = CC.CustomerID
	INNER JOIN ServiceOfferings.ServiceOfferings SO
		ON SO.ServiceOfferingID = CSO.CustomerServiceOfferingID
	WHERE 	
		C.EndDate IS NULL AND CC.EndDate IS NULL AND CSO.EndDate IS NULL AND SO.EndDate IS NULL
		AND
		((@apiKey = '' OR CC.ClientAPIKey = @apiKey)
		OR
		(@CustomerID = 0 OR  CC.CustomerID = @CustomerID))

END
