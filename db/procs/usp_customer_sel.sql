USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_customer_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_customer_sel')
BEGIN
    DROP PROCEDURE [Customer].[usp_customer_sel]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_customer_sel]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-04-23  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [Customer].[usp_customer_sel]
    (
    @CustomerID int
)
AS
BEGIN

    SET NOCOUNT ON;

    SELECT [CustomerID]
		  , [Name]
		  , [InternalCustomerID]
		  , [StartDate]
		  , [EndDate]
		  , [InsertedDate]
		  , [RowVersion]
    FROM [Customer].[Customer]
    WHERE 
		CustomerID = @CustomerID
        AND EndDate IS NULL

END

GO