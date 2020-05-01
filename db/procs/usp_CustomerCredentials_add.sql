USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_CustomerCredentials_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_CustomerCredentials_add')
BEGIN
    DROP PROCEDURE [Customer].[usp_CustomerCredentials_add]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_CustomerCredentials_add]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-04-23  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [Customer].[usp_CustomerCredentials_add]
    (
    @ClientAPIKey nvarchar(100),
    @Secret nvarchar(100),
    @CustomerID int,
    @startDate datetime,
    @endDate datetime,
    @UnitTestBypass bit=0
)
AS
BEGIN

    SET NOCOUNT ON;

    IF @@TRANCOUNT = 0
	BEGIN
        RAISERROR('Procedure must be called within a transaction', 16,1);
        RETURN;
    END

    IF EXISTS (SELECT 1
    FROM [Customer].[CustomerCredentials]
    WHERE @ClientAPIKey = ClientAPIKey)
	BEGIN
        RAISERROR('Client API key already exists', 16,1);
        RETURN;
    END


    IF EXISTS (SELECT 1
    FROM [Customer].[Customer]
    WHERE @CustomerID = CustomerID)
	BEGIN

        INSERT INTO [Customer].[CustomerCredentials]
            ([ClientAPIKey]
            ,[Secret]
            ,[CustomerID]
            ,[StartDate]
            ,[EndDate]
            ,[RowInserted])
        VALUES
            (@ClientAPIKey,
                @Secret,
                @CustomerID,
                @startDate,
                @endDate,
                GETUTCDATE())
    END
	ELSE
	BEGIN
        RAISERROR('CustomerID does not exists',16,1);
        RETURN;
    END

END

GO